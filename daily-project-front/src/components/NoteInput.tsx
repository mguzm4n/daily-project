import { useState } from 'react';
import useTextArea from '../hooks/useTextArea';
import { ErrorHolder } from '../utils/validations';

const NoteInput = () => {
  const [inputs, setInputs] = useState({
    title: "",
    description: ""
  });

  const [errHolder, setErrHolder] = useState(new ErrorHolder());

  const setInputValues = (key: string, val: string) => {
    setInputs(v => ({ ...v, [key]: val }))
  }

  const checkTitleInput = (input: string) => {
    errHolder.addValidation(input.trim().length == 0, "Título", "No puede ser vacío");
    errHolder.addValidation(input.length > 6, "Título", "No puede contener más de 32 carácteres");

    setErrHolder(errHolder.clone())
    setInputValues("title", input);
  }

  const checkDescriptionInput = (input: string) => {
    errHolder.addValidation(input.trim().length == 0, "Descripción", "No puede ser vacío");
    errHolder.addValidation(input.includes("aaa"), "Descripción", "No puede contener 3 As seguidas");

    setErrHolder(errHolder.clone())
    setInputValues("description", input);
  }

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

    <input value={inputs.title} onChange={e => checkTitleInput(e.target.value)} />
    <input value={inputs.description} onChange={e => checkDescriptionInput(e.target.value)} />
    <div className="bg-slate-400">
      {
        errHolder.getErrorsList().map(({ field, msgs }) => (<div key={field}>
          { field }:
          { msgs.map(msg => <div key={msg}>{ msg }</div>)}
          <br />
        </div>))
      }
    </div>
  </div>)
};

export default NoteInput;
