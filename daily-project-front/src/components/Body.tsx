
import { useQuery } from 'react-query';
import NoteInput from './NoteInput';
import { showUserNotes } from '../services/user';

function Body() {
  const { isError: notesError, isLoading: notesLoading, data: notes } = useQuery({
    queryKey: ["user", 1, "notes"],
    queryFn: () => showUserNotes(1)
  });

  return (
    <div className="w-screen flex flex-col justify-center items-center gap-4 dark:text-gray-200">
      <div className="w-[90%] flex flex-col">
        <NoteInput />
        <div>
          <p>Hoy</p>
          { notesLoading && "Cargando sus notas..." }
          { notesError && "Hubo un error trayendo las notas..."}
          { (!notesError && notes) && <div>
            
          </div>}
        </div>
        <div>
          <p>Ayer</p>
        </div>
      </div>
    </div>
  )
}

export default Body;
