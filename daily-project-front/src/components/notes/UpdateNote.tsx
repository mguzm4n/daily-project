import { useMutation } from "react-query";
import { Note, updateNote } from "../../services/notes";
import Modal from "../Modal";
import { useState } from "react";
interface UpdateNoteProps {
  note: Note;
  updatingNote: boolean;
  setUpdatingNote: (state: boolean) => void;
}

const UpdateNote = ({ note, updatingNote, setUpdatingNote }: UpdateNoteProps) => {
  const [content, setContent] = useState(note.content);

  const { mutate: updateNoteRequest, isLoading, isSuccess } = useMutation({
    mutationFn: ({ content}: { content: string}) => updateNote(note.id, content)
  });

  const onNoteUpdate = () => {
    if (note.content === content.trim()) return;
    updateNoteRequest({ content });
  }

  const sameContent = note.content === content.trim();
  console.log(sameContent);
  return (
    <Modal isVisible={updatingNote} setIsVisible={setUpdatingNote}>
      <div className="text-black bg-white shadow-sm px-2 py-4">
        { content }
        <div className="flex justify-between">
          <button className="underline bold text-sm"
            onClick={() => setUpdatingNote(false)}>
            Volver
          </button>
          <button className="disabled:opacity-70 bg-blue-500 text-white font-bold text-sm px-2 py-1 rounded-xl"
            disabled={(isLoading && isSuccess) || sameContent}
            onClick={onNoteUpdate}>
            Guardar
          </button>
        </div>
      </div>
    </Modal>
  )
};

export default UpdateNote;
