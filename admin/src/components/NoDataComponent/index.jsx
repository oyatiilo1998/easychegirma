import "./style.scss"
import InboxIcon from '@mui/icons-material/Inbox';
import { experimentalStyled, Typography } from "@mui/material";

const StyledTitle = experimentalStyled(Typography)(({theme }) => ({
  color: theme.palette.grey[500]
}))

const NoDataComponent = ({title = "Список пуст", isVisible}) => {
  if(!isVisible) return null

  return (
    <div className="NoDataComponent" >
      <div className="block">
        <div className="icon">
          <InboxIcon color="disabled" style={{fontSize: "50"}} />
        </div>
        <StyledTitle variant="body1"  >{title}</StyledTitle>
      </div>
    </div>
  )
}

export default NoDataComponent
