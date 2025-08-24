import { Node } from "@/types/proto2/model/markdown_pb";
import Renderer from "./Renderer";

interface Props {
  symbol: string;
  children: Node[];
}

const Italic: React.FC<Props> = ({ children }: Props) => {
  return (
    <em>
      {children.map((child, index) => (
        <Renderer key={index} index={index.toString()} node={child} />
      ))}
    </em>
  );
};

export default Italic;
