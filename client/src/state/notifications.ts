import { atom } from "recoil";

export default interface Notification {
    type: 'error' | 'success';
    text: string;
    title: string;
    id?: string;
}

export const notificationsState = atom<Notification[]>({
    key: 'notifications',
    default: []
})

