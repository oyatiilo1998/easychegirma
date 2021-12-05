import { Checkbox } from "@mui/material"
import { useMemo } from "react"
import generateRandomNumber from "../../utils/genereateRandomNumber"

const FCheckbox = ({ formik, name, label }) => {

  const randomId = useMemo(() => {
    return generateRandomNumber()
  }, [])

  return (
    <div>
      <Checkbox
        id={`checkbox-${randomId}`}
        style={{ transform: 'translatey(-1px)' }}
      />
      <label htmlFor={`checkbox-${randomId}`}>{label}</label>
    </div>
  )
}

export default FCheckbox
