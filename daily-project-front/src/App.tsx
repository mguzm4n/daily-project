
import './App.css';
import NoteInput from './components/NoteInput';

function App() {

  return (<>
    <div className="mb-5 shadow w-full flex bg-gray-100 items-center px-4 py-2.5">
      <p className="font-[Signika] font-semibold italic">Daily Project</p>
    </div>
    <div className="w-screen flex flex-col justify-center items-center gap-4">
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
  </>)
}

export default App
