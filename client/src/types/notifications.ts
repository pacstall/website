export default interface Notification {
    type: 'error' | 'success' | 'info' | 'warning';
    text: string;
    title: string;
    id?: string;
}