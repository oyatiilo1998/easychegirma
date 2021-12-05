import { DatePicker, LocalizationProvider } from "@mui/lab"
import { TextField } from "@mui/material"
import DateAdapter from "@mui/lab/AdapterMoment"
import moment from "moment"
import "moment/locale/ru";

moment.locale("ru");

const CDatePicker = ({
  value,
  onChange = () => {},
  label,
  width,
  error,
  ...props
}) => {
  return (
    <LocalizationProvider dateAdapter={DateAdapter}>
      <DatePicker
        inputFormat="YYYY-MM-DD"
        onChange={(val) => onChange(val ? val?.format("YYYY-MM-DD") : null)}
        value={value}
        clearable
        mask="____.__.__"
        {...props}
        renderInput={(params) => (
          <TextField
            helperText={error ? error.message : " "}
            style={{ width }}
            size="small"
            {...params}
            error={error}
            label={label}
          />
        )}
      />
    </LocalizationProvider>
  )
}

export default CDatePicker
