
import NoteInput from './NoteInput';

function Body() {
  return (
    <div className="w-screen flex flex-col justify-center items-center gap-4 dark:text-gray-200">
      <div className="w-[90%] flex flex-col">
        <NoteInput />
        <div>
          <p>Hoy</p>
        </div>
        <div>
          <p>Ayer</p>
        </div>
      </div>
    </div>
  )
}

export default Body;
