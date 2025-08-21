import { uniqueId } from "lodash-es";
import { makeAutoObservable } from "mobx";
import { memoServiceClient as memoServiceClientV2 } from "@/grpc";
import { ListMemosRequest as ListMemosRequestV2, CreateMemoRequest as CreateMemoRequestV2 } from "@/types/proto2/api/v2/memo_pb";
import { Memo as MemoV2 } from "@/types/proto2/model/memo_pb";
import { FieldMask } from "@/types/proto/google/protobuf/field_mask";

class LocalState {
  stateId: string = uniqueId();
  memoMapByName: Record<string, MemoV2> = {};
  currentRequest: AbortController | null = null;

  constructor() {
    makeAutoObservable(this);
  }

  setPartial(partial: Partial<LocalState>) {
    Object.assign(this, partial);
  }

  updateStateId() {
    this.stateId = uniqueId();
  }

  get memos() {
    return Object.values(this.memoMapByName);
  }

  get size() {
    return Object.keys(this.memoMapByName).length;
  }
}

const memoStore = (() => {
  const state = new LocalState();

  const fetchMemos = async (request: Partial<ListMemosRequestV2>) => {
    if (state.currentRequest) {
      state.currentRequest.abort();
    }

    const controller = new AbortController();
    state.setPartial({ currentRequest: controller });

    try {
      const { memos, nextPageToken } = await memoServiceClientV2.listMemos(request as ListMemosRequestV2, { signal: controller.signal });
      if (!controller.signal.aborted) {
        const memoMap = request.pageToken ? { ...state.memoMapByName } : {};
        for (const memo of memos) {
          memoMap[memo.name] = memo;
        }
        state.setPartial({
          stateId: uniqueId(),
          memoMapByName: memoMap,
        });
        return { memos, nextPageToken };
      }
    } catch (error: any) {
      if (error.name === "AbortError") {
        return;
      }
      throw error;
    } finally {
      if (state.currentRequest === controller) {
        state.setPartial({ currentRequest: null });
      }
    }
  };

  const getOrFetchMemoByName = async (name: string, options?: { skipCache?: boolean; skipStore?: boolean }) => {
    const memoCache = state.memoMapByName[name];
    if (memoCache && !options?.skipCache) {
      return memoCache;
    }

    const memo = await memoServiceClientV2.getMemo({
      name,
    });

    if (!options?.skipStore) {
      const memoMap = { ...state.memoMapByName };
      memoMap[name] = memo;
      state.setPartial({
        stateId: uniqueId(),
        memoMapByName: memoMap,
      });
    }

    return memo;
  };

  const getMemoByName = (name: string) => {
    return state.memoMapByName[name];
  };

  const createMemo = async (request: CreateMemoRequestV2) => {
    const memo = await memoServiceClientV2.createMemo(request);
    const memoMap = { ...state.memoMapByName };
    memoMap[memo.name] = memo;
    state.setPartial({
      stateId: uniqueId(),
      memoMapByName: memoMap,
    });
    return memo;
  };

  const updateMemo = async (update: Partial<MemoV2>, updateMask: string[]) => {
    const memo = await memoServiceClientV2.updateMemo({
      memo: update as MemoV2,
      updateMask: FieldMask.create({ paths: updateMask }),
    });

    const memoMap = { ...state.memoMapByName };
    memoMap[memo.name] = memo;
    state.setPartial({
      stateId: uniqueId(),
      memoMapByName: memoMap,
    });
    return memo;
  };

  const deleteMemo = async (name: string) => {
    await memoServiceClientV2.deleteMemo({
      name,
    });

    const memoMap = { ...state.memoMapByName };
    delete memoMap[name];
    state.setPartial({
      stateId: uniqueId(),
      memoMapByName: memoMap,
    });
  };

  return {
    state,
    fetchMemos,
    getOrFetchMemoByName,
    getMemoByName,
    createMemo,
    updateMemo,
    deleteMemo,
  };
})();

export default memoStore;
