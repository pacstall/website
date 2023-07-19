import { NumericDisplayHandler } from '../../hooks/useNumericDisplay'

export const arabic: NumericDisplayHandler = value => value.toString(10)
