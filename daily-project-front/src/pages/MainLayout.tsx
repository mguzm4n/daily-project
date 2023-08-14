import { Outlet } from "react-router-dom";
//import useAuthCheck from "../hooks/useAuthCheck";
import { useEffect, useState } from "react";
import { BsSun, BsMoon } from 'react-icons/bs';
import { MODES, getThemeFromLocal, getThemeFromMedia, setClassnameThemeMode } from "../utils/theme";

function ThemeSelector() {
  const [mode, setMode] = useState("");
  const isLight = mode === MODES.Light;

  // 
  useEffect(() => {
    const themeFromLocal = getThemeFromLocal();
    if (themeFromLocal) {
      setMode(themeFromLocal);
      return;
    }; // Get and set theme from local info.

    const themeFromMedia = getThemeFromMedia();
    setMode(themeFromMedia);
  }, []);

  const switchMode = () => {
    setClassnameThemeMode();
    setMode(isLight ? MODES.Dark : MODES.Light);
  }

  return(<div 
    onClick={switchMode}
    className="cursor-pointer flex items-center gap-1 text-black dark:text-gray-100">
    <p className="text-sm">
      Go {isLight ? MODES.Dark : MODES.Light}!
    </p>
    {isLight
        ? <BsSun />
        : <BsMoon className="text-sm -rotate-[5deg]" />}
  </div>);
}

const MainLayout = () => {
  // const authQuery = useAuthCheck();
  return (<>
    <div className="mb-5 shadow w-full flex justify-between bg-gray-100 items-center px-4 py-2.5 dark:bg-gray-600">
      <p className="dark:text-gray-200 font-[Signika] font-semibold italic">Daily Project</p>
      <ThemeSelector />
    </div>
    <Outlet />
  </>)
};

export default MainLayout;