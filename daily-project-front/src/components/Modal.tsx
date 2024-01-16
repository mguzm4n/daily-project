import { createPortal } from "react-dom";
import { ReactNode, useEffect } from "react";

interface ModalProps {
  children: ReactNode;
  isVisible: boolean;
  setIsVisible?: (visibility: boolean) => void;
}

const Modal = ({ children, isVisible, setIsVisible }: ModalProps)  => {

  const onCtrlPress = (e: KeyboardEvent, setIsVisibleFn: (visibility: boolean) => void) => {
    if (e.key == "Escape") {
      setIsVisibleFn(false);
    }
  };

  useEffect(() => {
    if (!setIsVisible) return;

    const cb = (e: KeyboardEvent) => onCtrlPress(e, setIsVisible);
    window.addEventListener('keydown', cb);
    return () => window.removeEventListener('keydown', cb);
  }, [setIsVisible]);
  
  if (!isVisible) return null;

  return createPortal(
    <div className="absolute left-0 top-0 w-full h-screen bg-slate-400 bg-opacity-20 flex items-center justify-center">
    { children }
    </div>,
    document.getElementById("modal")!,
    "modal"
  )
};

export default Modal;
