import { Typography } from "@mui/material"
import "./style.scss"

const Header = ({title, children}) => {
  return (
    <div className="Header" >
      <Typography variant="h4" color="secondary" >{title}</Typography>
      <div className="right-side">
        <div style={{marginLeft: "10px", display: "flex", gridGap: "10px"}}>
          {children}
        </div>
      </div>
    </div>
  )
}

export default Header
