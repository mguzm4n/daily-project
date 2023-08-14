import { useQuery } from "react-query";
import { getAuth } from "../services/auth";
import useAuth from "./useAuth";

export default function useAuthCheck() {
  const { setUser } = useAuth();
  return useQuery('auth', {
    queryFn: getAuth,
    onSuccess: (data) => {
      console.log('User is: ', data);
      setUser(data);
    },
    onError: (err) => {
      console.log('Auth Error: ', err);
      setUser(undefined);
    },
    retry: false,
    staleTime: 1000*60*30, // time for data till it's stale or not fresh
    cacheTime: 1000*60*35 // 30 min
  });
}