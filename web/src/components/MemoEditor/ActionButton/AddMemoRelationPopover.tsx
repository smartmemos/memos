import { Autocomplete, AutocompleteOption, Chip } from "@mui/joy";
import { Button, Checkbox } from "@usememos/mui";
import { uniqBy } from "lodash-es";
import { LinkIcon } from "lucide-react";
import React, { useContext, useState } from "react";
import { toast } from "react-hot-toast";
import useDebounce from "react-use/lib/useDebounce";
import { Popover, PopoverContent, PopoverTrigger } from "@/components/ui/Popover";
import { memoServiceClient } from "@/grpc";
import { DEFAULT_LIST_MEMOS_PAGE_SIZE } from "@/helpers/consts";
import useCurrentUser from "@/hooks/useCurrentUser";
import { extractMemoIdFromName } from "@/store/v1";
import { Memo, MemoRelation_Memo, MemoRelation_Type } from "@/types/proto/api/v1/memo_service";
import { useTranslate } from "@/utils/i18n";
import { EditorRefActions } from "../Editor";
import { MemoEditorContext } from "../types";

interface Props {
  editorRef: React.RefObject<EditorRefActions>;
}

const AddMemoRelationPopover = (props: Props) => {
  const { editorRef } = props;
  const t = useTranslate();
  const context = useContext(MemoEditorContext);
  const user = useCurrentUser();
  const [searchText, setSearchText] = useState<string>("");
  const [isFetching, setIsFetching] = useState<boolean>(true);
  const [fetchedMemos, setFetchedMemos] = useState<Memo[]>([]);
  const [selectedMemos, setSelectedMemos] = useState<Memo[]>([]);
  const [embedded, setEmbedded] = useState<boolean>(false);
  const [popoverOpen, setPopoverOpen] = useState<boolean>(false);

  const filteredMemos = fetchedMemos.filter(
    (memo) =>
      !selectedMemos.includes(memo) &&
      memo.name !== context.memoName &&
      !context.relationList.some((relation) => relation.relatedMemo?.name === memo.name),
  );

  useDebounce(
    async () => {
      if (!popoverOpen) return;

      setIsFetching(true);
      try {
        const conditions = [];
        if (searchText) {
          conditions.push(`content_search == [${JSON.stringify(searchText)}]`);
        }
        const { memos } = await memoServiceClient.listMemos({
          parent: user.name,
          pageSize: DEFAULT_LIST_MEMOS_PAGE_SIZE,
          oldFilter: conditions.length > 0 ? conditions.join(" && ") : undefined,
        });
        setFetchedMemos(memos);
      } catch (error: any) {
        toast.error(error.details);
        console.error(error);
      }
      setIsFetching(false);
    },
    300,
    [popoverOpen, searchText],
  );

  const getHighlightedContent = (content: string) => {
    const index = content.toLowerCase().indexOf(searchText.toLowerCase());
    if (index === -1) {
      return content;
    }
    let before = content.slice(0, index);
    if (before.length > 20) {
      before = "..." + before.slice(before.length - 20);
    }
    const highlighted = content.slice(index, index + searchText.length);
    let after = content.slice(index + searchText.length);
    if (after.length > 20) {
      after = after.slice(0, 20) + "...";
    }

    return (
      <>
        {before}
        <mark className="font-medium">{highlighted}</mark>
        {after}
      </>
    );
  };

  const addMemoRelations = async () => {
    // If embedded mode is enabled, embed the memo instead of creating a relation.
    if (embedded) {
      if (!editorRef.current) {
        toast.error(t("message.failed-to-embed-memo"));
        return;
      }

      const cursorPosition = editorRef.current.getCursorPosition();
      const prevValue = editorRef.current.getContent().slice(0, cursorPosition);
      if (prevValue !== "" && !prevValue.endsWith("\n")) {
        editorRef.current.insertText("\n");
      }
      for (const memo of selectedMemos) {
        editorRef.current.insertText(`![[memos/${extractMemoIdFromName(memo.name)}]]\n`);
      }
      setTimeout(() => {
        editorRef.current?.scrollToCursor();
        editorRef.current?.focus();
      });
    } else {
      context.setRelationList(
        uniqBy(
          [
            ...selectedMemos.map((memo) => ({
              memo: MemoRelation_Memo.fromPartial({ name: memo.name }),
              relatedMemo: MemoRelation_Memo.fromPartial({ name: memo.name }),
              type: MemoRelation_Type.REFERENCE,
            })),
            ...context.relationList,
          ].filter((relation) => relation.relatedMemo !== context.memoName),
          "relatedMemo",
        ),
      );
    }
    setSelectedMemos([]);
    setPopoverOpen(false);
  };

  return (
    <Popover open={popoverOpen} onOpenChange={setPopoverOpen}>
      <PopoverTrigger className="w-9 relative">
        <Button className="flex items-center justify-center" size="sm" variant="plain" asChild>
          <LinkIcon className="w-5 h-5 mx-auto p-0" />
        </Button>
      </PopoverTrigger>
      <PopoverContent align="center">
        <div className="w-[16rem] flex flex-col justify-start items-start">
          <Autocomplete
            className="w-full"
            size="md"
            clearOnBlur
            disableClearable
            placeholder={t("reference.search-placeholder")}
            noOptionsText={t("reference.no-memos-found")}
            options={filteredMemos}
            loading={isFetching}
            inputValue={searchText}
            value={selectedMemos}
            multiple
            onInputChange={(_, value) => setSearchText(value.trim())}
            getOptionKey={(memo) => memo.name}
            getOptionLabel={(memo) => memo.content}
            isOptionEqualToValue={(memo, value) => memo.name === value.name}
            renderOption={(props, memo) => (
              <AutocompleteOption {...props} key={memo.name}>
                <div className="w-full flex flex-col justify-start items-start">
                  <p className="text-xs text-gray-400 select-none">{memo.displayTime?.toLocaleString()}</p>
                  <p className="mt-0.5 text-sm leading-5 line-clamp-2">{searchText ? getHighlightedContent(memo.content) : memo.snippet}</p>
                </div>
              </AutocompleteOption>
            )}
            renderTags={(memos) =>
              memos.map((memo) => (
                <Chip key={memo.name} className="!max-w-full !rounded" variant="outlined" color="neutral">
                  <div className="w-full flex flex-col justify-start items-start">
                    <p className="text-xs text-gray-400 select-none">{memo.displayTime?.toLocaleString()}</p>
                    <span className="w-full text-sm leading-5 truncate">{memo.content}</span>
                  </div>
                </Chip>
              ))
            }
            onChange={(_, value) => setSelectedMemos(value)}
          />
          <div className="mt-2 w-full flex flex-row justify-end items-center gap-2">
            <Checkbox size="sm" label={"Embed"} checked={embedded} onChange={(e) => setEmbedded(e.target.checked)} />
            <Button size="sm" color="primary" onClick={addMemoRelations} disabled={selectedMemos.length === 0}>
              {t("common.add")}
            </Button>
          </div>
        </div>
      </PopoverContent>
    </Popover>
  );
};

export default AddMemoRelationPopover;
