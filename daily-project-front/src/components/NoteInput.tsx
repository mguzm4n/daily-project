import { useState } from 'react';
import useTextArea from '../hooks/useTextArea';
import { ErrorHolder } from '../utils/validations';
import { useMutation } from 'react-query';
import { createNote } from '../services/notes';

const NoteInput = () => {
  const [uiErrorState, setUiErrorState] = useState({
    omitErrors: false, alertErrors: false
  });
  const [errHolder, setErrHolder] = useState(new ErrorHolder());

  const { body, setBody, textAreaRef } = useTextArea();
  const { mutate: sendNote, isLoading: isCreatingNote } = useMutation({
    mutationFn: ({ content }: { content: string }) => createNote(content),
    onSuccess: () => {
      setBody("");
    }
  });


  const setUiErrorsState = (key: 'omitErrors' | 'alertErrors', value: boolean) => {
    setUiErrorState( v => ({ ...v, [key]: value}));
  }
  const onSaveNoteClick = () => {
    setUiErrorsState('omitErrors', false);

    if (!errHolder.passedValidations()) {
      // remember there are errors on some state...
      setUiErrorsState('alertErrors', true);
      setTimeout(() => {
        setUiErrorsState('alertErrors', false);
      }, 805);
      return;
    }
    sendNote({ content: body });
  };

  const setBodyContent = (input: string) => {
    errHolder.addValidation(input.trim().length == 0, "Contenido", "No puede ser vacío");

    setErrHolder(errHolder.clone());
    setBody(input);
  };

  return (<div className="bg-white rounded-md shadow px-8 py-4 w-full flex flex-col gap-2
    dark:bg-zinc-700">
    <textarea placeholder="Registra tu momento..." 
      className="overflow-hidden resize-none rounded-3xl pt-3 pb-9 px-4 border-2 border-gray-300
        dark:focus:border-pink-500 dark:border-zinc-600 dark:text-white dark:placeholder-white dark:bg-zinc-500
        focus:outline-none focus:border-pink-500"
      ref={textAreaRef} name="content" id="content"
      value={body} onChange={e => setBodyContent(e.target.value)} 
    />
    <button disabled={isCreatingNote}
      onClick={onSaveNoteClick}
      className="focus:outline-pink-600 hover:ring-2 hover:ring-pink-500 hover:bg-transparent hover:text-black 
      rounded-3xl self-end w-1/3 bg-pink-500 text-white text-xs font-medium tracking-wide py-3 uppercase
      dark:hover:text-white">
        Guardar
    </button>
    {(!uiErrorState.omitErrors && errHolder.hasErrors()) &&  
    <div className={`rounded px-4 py-2 text-sm dark:text-gray-300 dark:bg-zinc-800 bg-slate-200 text-gray-800 
      ${uiErrorState.alertErrors ? "animate-[wiggle_400ms_ease-in-out_infinite]" : ""}`}>
      {
        errHolder.getErrorsList().map(({ field, msgs }) => (<div key={field}>
          { field }:
          { msgs.map(msg => <div key={msg}>{ msg }</div>)}
          <br />
        </div>))
      }
      <button onClick={() => setUiErrorsState('omitErrors', true)}>Ok</button>
    </div>}
  </div>)
};

export default NoteInput;
