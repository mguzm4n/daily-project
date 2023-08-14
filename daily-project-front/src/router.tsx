import { createBrowserRouter } from "react-router-dom";
import MainLayout from "./pages/MainLayout";
import Home from "./pages/Home";
import RequireAuth from "./pages/RequireAuth";
import Body from "./components/Body";

const protectedRoutes = {
  element: <RequireAuth />,
  children: [
    { path: "home", element: <Body />}
  ]
}
const router = createBrowserRouter([{
  path: "/",
  element: <MainLayout />,
  children: [
    protectedRoutes
  ]
}]);
 
 export default router;