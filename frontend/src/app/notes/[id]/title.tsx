import TextAreAutosize from "react-textarea-autosize";
import {useState} from "react";


interface TitleProps {
  initialValue: string;
  onChange: (title: string) => void;
}

const Title :  React.FC<TitleProps> = ({initialValue, onChange}) => {
  const [title, setTitle] = useState<string>(initialValue || "");
  return (
    <TextAreAutosize
      value={title}
      className="w-full resize-none appearance-none overflow-hidden bg-transparent text-5xl font-bold focus:outline-none px-0"
      onChange={async (e) => {
        setTitle(e.target.value)
        onChange(title)
      }}
    />
  );
}

export default Title;
