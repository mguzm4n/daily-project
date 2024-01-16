import { formatTweetDate } from "../../utils/time";
import { FiEdit2 } from 'react-icons/fi'; 
import { useState } from "react";
import UpdateNote from "./UpdateNote";
import { Note } from "../../services/notes";

const NotePreview = ({ note }: { note: Note }) => {
  const [updateNote, setUpdateNote] = useState(false);

  return (
    <div className="bg-white rounded-md shadow px-8 py-4 w-full flex flex-col gap-2
    dark:bg-zinc-700"> 
      <button onClick={() => setUpdateNote(true)}
        className="text-xs font-medium flex items-center gap-2">
        <FiEdit2 />
        Editar
      </button>
      { note.content }
      <p className="text-xs">
        <p>Capturado a las {formatTweetDate(note.createdAt.toUTCString())}</p>
        <p>
          Última vez 
          <button className="px-1 font-medium"
            title="Ver historial de modificaciones">actualizado</button>
          a las
        </p>
      </p>
    <UpdateNote note={note}
      updatingNote={updateNote} setUpdatingNote={setUpdateNote} />
    </div>
  )
}

export default NotePreview;