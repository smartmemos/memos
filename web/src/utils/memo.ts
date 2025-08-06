// import { Visibility } from "@/types/proto/api/v1/memo_service";
import { Visibility as VisibilityV2 } from "@/types/proto2/model/common_pb";

export const convertVisibilityFromString = (visibility: string) => {
  switch (visibility) {
    case "PUBLIC":
      return VisibilityV2.PUBLIC;
    case "PROTECTED":
      return VisibilityV2.PROTECTED;
    case "PRIVATE":
      return VisibilityV2.PRIVATE;
    default:
      return VisibilityV2.PUBLIC;
  }
};

export const convertVisibilityToString = (visibility: VisibilityV2) => {
  switch (visibility) {
    case VisibilityV2.PUBLIC:
      return "PUBLIC";
    case VisibilityV2.PROTECTED:
      return "PROTECTED";
    case VisibilityV2.PRIVATE:
      return "PRIVATE";
    default:
      return "PRIVATE";
  }
};
