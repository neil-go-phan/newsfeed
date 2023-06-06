import axiosProtectedAPI from './axiosProtectedAPI';

export async function checkAuth(): Promise<boolean> {
  const res: any = await axiosProtectedAPI.get('/auth/check-auth').then((res) => res);
  if (res.data) {
    return true
  }
  return false;
}
