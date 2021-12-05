import { LocalizationProvider } from "@mui/lab";
import { useSelector } from "react-redux";
import ThemeConfig from "./theme";
// import DateAdapter from '@mui/lab/AdapterMoment';
import AdapterDateFns from '@mui/lab/AdapterDateFns';


const MaterialProvider = ({children}) => {
  const theme = "light"

  return (
    <div className={theme === "dark" && 'night-mode'} >
      <ThemeConfig >
        <LocalizationProvider dateAdapter={AdapterDateFns} >
          {children}
        </LocalizationProvider>
      </ThemeConfig>
    </div>
  );
};

export default MaterialProvider;
