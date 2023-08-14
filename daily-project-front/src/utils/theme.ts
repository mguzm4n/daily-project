
export const MODES = { Light: "light", Dark: "dark"};

export function toggleClasslist({ add, remove }: { add: string, remove: string}) {
    document.documentElement.classList.remove(remove);
    document.documentElement.classList.add(add);
}

export function setClassnameThemeMode() {
    if (document.documentElement.classList.contains(MODES.Light)) {
      toggleClasslist({ add: MODES.Dark, remove: MODES.Light });
      localStorage.setItem('mode', MODES.Dark);

    } else if (document.documentElement.classList.contains(MODES.Dark)) {
      toggleClasslist({ add: MODES.Light, remove: MODES.Dark });
      localStorage.setItem('mode', MODES.Light);
    }
}

export function getThemeFromLocal() {
  let modeOnLocalStorage = localStorage.getItem("mode");
  if (modeOnLocalStorage) {
    if (modeOnLocalStorage !== MODES.Dark || modeOnLocalStorage !== MODES.Light) {
      return null;
    }

    const isLightMode = modeOnLocalStorage === MODES.Light;
    if (isLightMode) {
      document.body.classList.remove(MODES.Dark);
      document.body.classList.add(MODES.Light);
    } else {
      document.body.classList.remove(MODES.Light);
      document.body.classList.add(MODES.Dark);
    }
  }

  return modeOnLocalStorage;
}

export function getThemeFromMedia() {
  if (window.matchMedia && window.matchMedia('(prefers-color-scheme: dark)').matches) {
    document.documentElement.classList.add(MODES.Dark);
    localStorage.setItem("mode", MODES.Dark);
    return MODES.Dark;
  } else {
    document.documentElement.classList.add(MODES.Light);
    localStorage.setItem("mode", MODES.Light);
    return MODES.Light;
  }
}