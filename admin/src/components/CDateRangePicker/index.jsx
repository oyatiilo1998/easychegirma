import { DesktopDateRangePicker, LocalizationProvider } from "@mui/lab"
import DateAdapter from "@mui/lab/AdapterMoment"
import { TextField } from "@mui/material"

const CDateRangePicker = ({ onChange = () => {}, value }) => {
  return (
    <LocalizationProvider dateAdapter={DateAdapter}>
      <DesktopDateRangePicker
        // label="Date desktop"
        inputFormat="DD-MM-YYYY"
        value={value}
        onChange={(val) => {
          onChange(val.map((el) => (el ? el.format("YYYY-MM-DD") : null)))
        }}
        renderInput={(startProps, endProps) => {
          return (
            <>
              <TextField
                size="small"
                style={{ width: "150px", marginRight: "20px" }}
                {...startProps}
                label="Boshlangich vaqt"
              />
              {/* <Box sx={{ mx: 2 }}> to </Box> */}
              <TextField
                size="small"
                style={{ width: "150px" }}
                {...endProps}
                label="Tugash vaqti"
              />
            </>
          )
        }}
      />
    </LocalizationProvider>
  )
}

export default CDateRangePicker
