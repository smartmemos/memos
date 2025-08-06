import { Globe2Icon, LockIcon, UsersIcon } from "lucide-react";
import { cn } from "@/lib/utils";
import { Visibility as VisibilityV2 } from "@/types/proto2/model/common_pb";

interface Props {
  visibility: VisibilityV2;
  className?: string;
}

const VisibilityIcon = (props: Props) => {
  const { className, visibility } = props;

  let VIcon = null;
  if (visibility === VisibilityV2.PRIVATE) {
    VIcon = LockIcon;
  } else if (visibility === VisibilityV2.PROTECTED) {
    VIcon = UsersIcon;
  } else if (visibility === VisibilityV2.PUBLIC) {
    VIcon = Globe2Icon;
  }
  if (!VIcon) {
    return null;
  }

  return <VIcon className={cn("w-4 h-auto text-muted-foreground", className)} />;
};

export default VisibilityIcon;
