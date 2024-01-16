import axios from './axios';
const url = "/v1/notes";

export type Note = {
    id: number,
    content: string,
    updatedAt: Date,
    createdAt: Date,
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

export async function updateNote(id: number, content: string) {
    return axios
        .put<{ note: Note }>(`${url}/${id}`, { content })
        .then(r => r.data.note);
}