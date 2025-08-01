import { Dropdown, Menu, MenuButton } from "@mui/joy";
import { SmilePlusIcon } from "lucide-react";
import { useRef, useState } from "react";
import useClickAway from "react-use/lib/useClickAway";
import { memoServiceClient } from "@/grpc";
import useCurrentUser from "@/hooks/useCurrentUser";
import { useMemoStore } from "@/store/v1";
import { workspaceStore } from "@/store/v2";
import { Memo } from "@/types/proto/api/v1/memo_service";
import { cn } from "@/utils";

interface Props {
  memo: Memo;
  className?: string;
}

const ReactionSelector = (props: Props) => {
  const { memo, className } = props;
  const currentUser = useCurrentUser();
  const memoStore = useMemoStore();
  const [open, setOpen] = useState(false);
  const containerRef = useRef<HTMLDivElement>(null);
  const workspaceMemoRelatedSetting = workspaceStore.state.memoRelatedSetting;

  useClickAway(containerRef, () => {
    setOpen(false);
  });

  const hasReacted = (reactionType: string) => {
    return memo.reactions.some((r) => r.reactionType === reactionType && r.creator === currentUser?.name);
  };

  const handleReactionClick = async (reactionType: string) => {
    try {
      if (hasReacted(reactionType)) {
        const reactions = memo.reactions.filter(
          (reaction) => reaction.reactionType === reactionType && reaction.creator === currentUser.name,
        );
        for (const reaction of reactions) {
          await memoServiceClient.deleteMemoReaction({ id: reaction.id });
        }
      } else {
        await memoServiceClient.upsertMemoReaction({
          name: memo.name,
          reaction: {
            contentId: memo.name,
            reactionType: reactionType,
          },
        });
      }
      await memoStore.getOrFetchMemoByName(memo.name, { skipCache: true });
    } catch (error) {
      // skip error.
    }
    setOpen(false);
  };

  return (
    <Dropdown open={open} onOpenChange={(_, isOpen) => setOpen(isOpen)}>
      <MenuButton slots={{ root: "div" }}>
        <span
          className={cn("h-7 w-7 flex justify-center items-center rounded-full border dark:border-zinc-700 hover:opacity-70", className)}
        >
          <SmilePlusIcon className="w-4 h-4 mx-auto text-gray-500 dark:text-gray-400" />
        </span>
      </MenuButton>
      <Menu className="relative" component="div" size="sm" placement="bottom-start">
        <div ref={containerRef}>
          <div className="flex flex-row flex-wrap py-0.5 px-2 h-auto gap-1 max-w-56">
            {workspaceMemoRelatedSetting.reactions.map((reactionType) => {
              return (
                <span
                  key={reactionType}
                  className={cn(
                    "inline-flex w-auto text-base cursor-pointer rounded px-1 text-gray-500 dark:text-gray-400 hover:opacity-80",
                    hasReacted(reactionType) && "bg-blue-100 dark:bg-zinc-800",
                  )}
                  onClick={() => handleReactionClick(reactionType)}
                >
                  {reactionType}
                </span>
              );
            })}
          </div>
        </div>
      </Menu>
    </Dropdown>
  );
};

export default ReactionSelector;
