import { FC } from "react"
import { Link } from "react-router-dom"

const PageLink: FC<{ page: number, active?: boolean, disabled?: boolean }> = ({ page, active, disabled }) => (
    <Link className="no-underline" to={disabled ? '#' : location.search.replace(/page=[0-9]*/, `page=${page}`)}>
        <li className={`text-xs mx-1 px-3 py-2 bg-gray-200 ${active ? 'text-gray-700' : disabled ? 'text-gray-400' : 'text-gray-500'} hover:bg-gray-700 hover:text-gray-200 rounded-lg`}>
            <span className="font-bold">{page || '0'}</span>
        </li>
    </Link>
)

const PageSequentialLink: FC<{ nextPage: number | string, text: string, disabled?: boolean }> = ({ nextPage, text, disabled }) => (
    <Link className="no-underline" to={disabled ? '#' : location.search.replace(/page=[0-9]*/, `page=${nextPage}`)}>
        <li className={`text-xs mx-1 px-3 py-2 bg-gray-200 ${disabled ? 'text-gray-400' : 'text-gray-700'} hover:bg-gray-700 hover:text-gray-200 rounded-lg`}>
            <span className="flex items-center">
                <span className="mx-1">{text}</span>
            </span>
        </li>
    </Link>
)

const Dots: FC = () => <span style={{ transform: 'translateY(38%)' }}>...</span>

const Pagination: FC<{ last: number, current: number }> = ({ last, current }) => {

    return (
        <ul className="flex">
            <PageSequentialLink text="previous" nextPage={current - 1} disabled={current === 0} />


            {current <= 1 ? (
                <>
                    {new Array(current + 1).fill(null).map((_, i) => <PageLink key={i} page={i} active={i === current} />)}
                </>
            ) : (
                <>
                    <PageLink page={0} />
                    <Dots />
                    <PageLink page={current - 1} />
                    <PageLink page={current} active />
                </>
            )}

            {current + 1 < last && <PageLink page={current + 1} />}
            {current + 2 < last ? (
                <>
                    <Dots />
                    <PageLink page={last} />
                </>
            ) : (current < last && <PageLink page={last} />)}

            <PageSequentialLink text="next" nextPage={current + 1} disabled={current === last} />
        </ul>
    )
}

export default Pagination