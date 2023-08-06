import { useState } from 'react';
import useTextArea from '../hooks/useTextArea';

const NoteInput = () => {
  
  const { body, setBody, textAreaRef } = useTextArea();
  return (<div className="bg-white rounded-md shadow px-8 py-4 w-full flex flex-col gap-2">
    <textarea placeholder="Registra tu momento..." 
      className="focus:outline-none focus:border-pink-500
        overflow-hidden resize-none rounded-3xl pt-3 pb-9 px-4 border-2 border-gray-300"
      ref={textAreaRef} name="content" id="content"
      value={body} onChange={e => setBody(e.target.value)} 
    />
    <button className="focus:outline-pink-600 hover:ring-2 hover:ring-pink-500 hover:bg-white hover:text-black 
      rounded-3xl self-end w-1/3 bg-pink-500 text-white text-xs font-medium tracking-wide py-3 uppercase">
        Guardar
    </button>
  </div>)
};

export default NoteInput;
