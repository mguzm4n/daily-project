
import { useQuery } from 'react-query';
import NoteInput from './NoteInput';
import { showUserNotes } from '../services/user';
import NotePreview from './notes/NotePreview';

function Body() {
  const { isError: notesError, isLoading: notesLoading, data: notes } = useQuery({
    queryKey: ["user", 1, "notes"],
    queryFn: () => showUserNotes(1)
  });

  const notesOkCheck = !notesLoading && !notesError && notes;

  return (
    <div className="w-screen flex flex-col justify-center items-center gap-4 dark:text-gray-200">
      <div className="w-[90%] flex flex-col">
        <NoteInput />
        <div>
          <p>Hoy</p>
          { notesLoading && "Cargando sus notas..." }
          { notesError && "Hubo un error trayendo las notas..."}
          { notesOkCheck && <div>
            
          </div>}
          <NotePreview note={ { id: 1, content: "hii!", createdAt: new Date(), updatedAt: new Date()} } />
        </div>
        <div>
          <p>Ayer</p>
        </div>
      </div>
    </div>
  )
}

export default Body;
