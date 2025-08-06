import VisibilityIcon from "@/components/VisibilityIcon";
import { Select, SelectContent, SelectItem, SelectTrigger, SelectValue } from "@/components/ui/select";
import { Visibility as VisibilityV2 } from "@/types/proto2/model/common_pb";
import { useTranslate } from "@/utils/i18n";

interface Props {
  value: VisibilityV2;
  onChange: (visibility: VisibilityV2) => void;
  onOpenChange?: (open: boolean) => void;
}

const VisibilitySelector = (props: Props) => {
  const { value, onChange } = props;
  const t = useTranslate();

  const visibilityOptions = [
    { value: VisibilityV2.PRIVATE, label: t("memo.visibility.private") },
    { value: VisibilityV2.PROTECTED, label: t("memo.visibility.protected") },
    { value: VisibilityV2.PUBLIC, label: t("memo.visibility.public") },
  ];

  const handleOpenChange = (open: boolean) => {
    if (props.onOpenChange) {
      props.onOpenChange(open);
    }
  };

  return (
    <Select value={value.toString()} onValueChange={onChange} onOpenChange={handleOpenChange}>
      <SelectTrigger size="xs" className="!bg-background">
        <SelectValue />
      </SelectTrigger>
      <SelectContent align="end">
        {visibilityOptions.map((option) => (
          <SelectItem key={option.value} value={option.value.toString()}>
            <VisibilityIcon className="size-3.5" visibility={option.value} />
            {option.label}
          </SelectItem>
        ))}
      </SelectContent>
    </Select>
  );
};

export default VisibilitySelector;
