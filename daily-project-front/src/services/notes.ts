import axios from './axios';
const url = "/v1/notes";

export type Note = {
    id: number,
    content: string,
    updated_at: Date,
    created_at: Date,
};

export async function showNote(id: number) {
    return axios
        .get<{ note: Note }>(`${url}/${id}`)
        .then(r => r.data.note);
}

export async function createNote(content: string) {
    return axios
        .post<{ note: Note }>(url, { content })
        .then(r => r.data.note);
}