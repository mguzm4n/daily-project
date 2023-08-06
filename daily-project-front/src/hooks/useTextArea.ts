import { useEffect, useRef, useState } from 'react';

const useTextArea = () => {
  const textAreaRef = useRef<HTMLTextAreaElement>(null);
  const [body, setBody] = useState("");

  useEffect(() => {
    if (textAreaRef.current == null) return;
    textAreaRef.current.style.height = "0px";
    const scrollHeight = textAreaRef.current.scrollHeight;
    textAreaRef.current.style.height = `${scrollHeight}px`;
  }, [body, textAreaRef]);

  return { 
    body,
    setBody, 
    textAreaRef 
  };
};

export default useTextArea;