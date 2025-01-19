import { NumericDisplayHandler } from '../../hooks/useNumericDisplay'

export const bengali: NumericDisplayHandler = value =>
    value.toString(10).replace(/\d/g, d => '০১২৩৪৫৬৭৮৯'[d as any])
