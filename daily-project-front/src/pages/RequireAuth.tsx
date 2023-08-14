import { ReactNode } from "react";
import { Navigate, Outlet } from "react-router-dom";
// import useAuthCheck from "../hooks/useAuthCheck";

const RequireAuth = ({ children } : { children?: ReactNode }) => {
  //const { isLoading, isError } = useAuthCheck();
  let isLoading = false;
  let isError = false;

  if (isLoading) {
    return <div>Cargando credenciales...</div>;
  }

  if (isError) {
    return <Navigate to="/sign-in" />;
  }
  
  return children 
    ? <>{ children }</> 
    : <Outlet />;
};

export default RequireAuth;