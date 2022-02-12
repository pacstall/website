import { FC } from "react";
import { useRecoilState } from "recoil";
import Notification, { notificationsState } from "../state/notifications";

const colors = {
    success: '#44C997',
    info: '#4086F4',
    error: '#EB4132'
}

export const SimpleNotification: FC<Notification> = ({ type, text, title }) => (
    <div className="flex notification">
        <div className="mx-3 my-3">
            <div className="bg-white rounded-lg border-gray-300 border p-3 shadow-lg">
                <div className="flex flex-row">
                    <div className="px-2 py-2">
                        <svg width="24" height="24" viewBox="0 0 1792 1792" fill={colors[type] || colors.info} xmlns="http://www.w3.org/2000/svg">
                            <path d="M1299 813l-422 422q-19 19-45 19t-45-19l-294-294q-19-19-19-45t19-45l102-102q19-19 45-19t45 19l147 147 275-275q19-19 45-19t45 19l102 102q19 19 19 45t-19 45zm141 83q0-148-73-273t-198-198-273-73-273 73-198 198-73 273 73 273 198 198 273 73 273-73 198-198 73-273zm224 0q0 209-103 385.5t-279.5 279.5-385.5 103-385.5-103-279.5-279.5-103-385.5 103-385.5 279.5-279.5 385.5-103 385.5 103 279.5 279.5 103 385.5z" />
                        </svg>
                    </div>
                    <div className="ml-2 mr-6">
                        <span className="text-sm font-semibold">{title}</span>
                        <span className="text-sm block text-gray-500">{text}</span>
                    </div>
                </div>
            </div>
        </div>
    </div>
)

const Notifications: FC = () => {
    const [notifications] = useRecoilState(notificationsState)

    return (
        <div style={{ position: 'absolute', top: '0', right: '0' }}>
            {notifications.map(n => <SimpleNotification key={n.id!} {...n} />)}
        </div>
    )
}

export default Notifications