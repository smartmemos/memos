import { observer } from "mobx-react-lite";
import { useMemo } from "react";
import MemoView from "@/components/MemoView";
import PagedMemoList from "@/components/PagedMemoList";
import useCurrentUser from "@/hooks/useCurrentUser";
import { viewStore, userStore, workspaceStore } from "@/store";
import { extractUserIdFromName } from "@/store/common";
import memoFilterStore from "@/store/memoFilter";
import { State as StateV2 } from "@/types/proto2/model/common_pb";
import { Memo as MemoV2 } from "@/types/proto2/model/memo_pb";
import { WorkspaceSetting_Key } from "@/types/proto/api/v1/workspace_service";

// Helper function to extract shortcut ID from resource name
// Format: users/{user}/shortcuts/{shortcut}
const getShortcutId = (name: string): string => {
  const parts = name.split("/");
  return parts.length === 4 ? parts[3] : "";
};

const Home = observer(() => {
  const user = useCurrentUser();
  const selectedShortcut = userStore.state.shortcuts.find((shortcut) => getShortcutId(shortcut.name) === memoFilterStore.shortcut);

  const memoFilter = useMemo(() => {
    const conditions = [`creator_id == ${extractUserIdFromName(user.name)}`];
    if (selectedShortcut?.filter) {
      conditions.push(selectedShortcut.filter);
    }
    for (const filter of memoFilterStore.filters) {
      if (filter.factor === "contentSearch") {
        conditions.push(`content.contains("${filter.value}")`);
      } else if (filter.factor === "tagSearch") {
        conditions.push(`tag in ["${filter.value}"]`);
      } else if (filter.factor === "pinned") {
        conditions.push(`pinned`);
      } else if (filter.factor === "property.hasLink") {
        conditions.push(`has_link`);
      } else if (filter.factor === "property.hasTaskList") {
        conditions.push(`has_task_list`);
      } else if (filter.factor === "property.hasCode") {
        conditions.push(`has_code`);
      } else if (filter.factor === "displayTime") {
        const displayWithUpdateTime = workspaceStore.getWorkspaceSettingByKey(WorkspaceSetting_Key.MEMO_RELATED).memoRelatedSetting
          ?.displayWithUpdateTime;
        const factor = displayWithUpdateTime ? "updated_ts" : "created_ts";
        const filterDate = new Date(filter.value);
        const filterUtcTimestamp = filterDate.getTime() + filterDate.getTimezoneOffset() * 60 * 1000;
        const timestampAfter = filterUtcTimestamp / 1000;
        conditions.push(`${factor} >= ${timestampAfter} && ${factor} < ${timestampAfter + 60 * 60 * 24}`);
      }
    }
    console.log("conditions", conditions);
    return conditions.length > 0 ? conditions.join(" && ") : undefined;
  }, [memoFilterStore.filters, selectedShortcut?.filter]);

  return (
    <div className="w-full min-h-full bg-background text-foreground">
      <PagedMemoList
        renderer={(memo: MemoV2) => <MemoView key={`${memo.name}-${memo.displayTime}`} memo={memo} showVisibility showPinned compact />}
        listSort={(memos: MemoV2[]) =>
          memos
            .filter((memo) => memo.state === StateV2.NORMAL)
            .sort((a, b) =>
              viewStore.state.orderByTimeAsc
                ? Number(a.displayTime?.seconds || 0) - Number(b.displayTime?.seconds || 0)
                : Number(b.displayTime?.seconds || 0) - Number(a.displayTime?.seconds || 0),
            )
        }
        orderBy={viewStore.state.orderByTimeAsc ? "display_time asc" : "display_time desc"}
        filter={memoFilter}
      />
    </div>
  );
});

export default Home;
