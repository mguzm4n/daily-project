import axios from 'axios';
import { Note } from './notes';

const url = "/v1/users";


export async function showUserNotes(id: number) {
    return axios
        .get<{ notes: Note[] }>(`${url}/${id}/notes`)
        .then(r => r.data.notes);
}