import PropTypes from "prop-types"
import { useEffect } from "react"
import { Link as RouterLink, useLocation, matchPath } from "react-router-dom"
// components
// import Logo from '../../components/Logo';
import Scrollbar from "../Scrollbar/index"

import MenuLinks from "./SidebarConfig"
import SidebarItem, { ListItemStyle } from "./SidebarItem"
import {
  Avatar,
  Drawer,
  Hidden,
  Link,
  List,
  ListSubheader,
  Typography,
  Box,
  experimentalStyled,
  ListItemIcon,
  ListItemText,
} from "@mui/material"
import Logo from "../Logo"
import ExitToAppIcon from '@mui/icons-material/ExitToApp';
import { useDispatch, useSelector } from "react-redux"
import { authActions } from "../../redux/slices/auth.slice"

// ----------------------------------------------------------------------

const DRAWER_WIDTH = 280

const RootStyle = experimentalStyled("div")(({ theme }) => ({
  [theme.breakpoints.up("lg")]: {
    flexShrink: 0,
    width: DRAWER_WIDTH,
  
  },
}))

const AccountStyle = experimentalStyled("div")(({ theme }) => ({
  display: "flex",
  alignItems: "center",
  padding: theme.spacing(2, 2.5),
  margin: theme.spacing(1, 2.5, 5),
  borderRadius: theme.shape.borderRadiusSm,
  backgroundColor: theme.palette.grey[500_12],
  color: '#fff'
}))

// ----------------------------------------------------------------------

function reduceChild({ array, item, pathname, level }) {
  const key = item.href + level

  if (item.items) {
    const match = matchPath(pathname, {
      path: item.href,
      exact: false,
    })

    array = [
      ...array,
      <SidebarItem
        key={key}
        level={level}
        icon={item.icon}
        info={item.info}
        href={item.href}
        title={item.title}
        open={Boolean(match)}
      >
        {renderSidebarItems({
          pathname,
          level: level + 1,
          items: item.items,
        })}
      </SidebarItem>,
    ]
  } else {
    array = [
      ...array,
      <SidebarItem
        key={key}
        level={level}
        href={item.href}
        icon={item.icon}
        info={item.info}
        title={item.title}
      />,
    ]
  }
  return array
}

function renderSidebarItems({ items, pathname, level = 0 }) {
  return (
    <List disablePadding >
      {items.reduce(
        (array, item) => reduceChild({ array, item, pathname, level }),
        []
      )}
    </List>
  )
}

Sidebar.propTypes = {
  isOpenSidebar: PropTypes.bool,
  onCloseSidebar: PropTypes.func,
}

export default function Sidebar({ isOpenSidebar, onCloseSidebar }) {
  const { pathname } = useLocation()
  const dispatch = useDispatch()
  const userLogin = useSelector(state => state.auth.userLogin)

  useEffect(() => {
    if (isOpenSidebar && onCloseSidebar) {
      onCloseSidebar()
    }
    // eslint-disable-next-line react-hooks/exhaustive-deps
  }, [pathname])

  const renderContent = (
    <Scrollbar >
      <Box sx={{ px: 2.5, py: 3 }} style={{backgroundColor: '#1E2F40'}} >
        <RouterLink to="/" style={{ textDecoration: "none" }}>
          <Logo size={'small'} />
        </RouterLink>
      </Box>

      <Link underline="none" component={RouterLink} to="#">
        <AccountStyle>
          <Avatar
            alt={userLogin?.charAt(0).toUpperCase()}
            src="/static/mock-images/avatars/avatar_default.jpg"
          />
          <Box sx={{ ml: 2 }}>
            <Typography style={{color: "#fff"}} variant="subtitle2" sx={{ color: "text.primary" }}>
              {userLogin?.charAt(0).toUpperCase() + userLogin?.slice(1)}
            </Typography>
          </Box>
        </AccountStyle>
      </Link>

      {MenuLinks.map((list) => (
        <List
          disablePadding
          key={list.subheader}
          subheader={
            <ListSubheader
              disableSticky
              disableGutters
              sx={{
                mt: 3,
                mb: 2,
                pl: 5,
                color: "text.primary",
                typography: "overline",
              }}
            >
              {list.subheader}
            </ListSubheader>
          }
        >
          {renderSidebarItems({
            items: list.items,
            pathname,
          })}
        </List>
      ))}

      <ListItemStyle
        button
        disableGutters
        style={{position: "absolute", bottom: "30px", left: 0}}
        onClick={() => dispatch(authActions.logout())}
      >
        <ListItemIcon><ExitToAppIcon /></ListItemIcon>
        <ListItemText disableTypography primary={"Выйти"} />
      </ListItemStyle>

    </Scrollbar>
  )

  return (
    <RootStyle>
      <Hidden lgUp>
        <Drawer
          open={isOpenSidebar}
          onClose={onCloseSidebar}
          PaperProps={{
            sx: { width: DRAWER_WIDTH },
          }}
        >
          {renderContent}
        </Drawer>
      </Hidden>
      <Hidden lgDown>
        <Drawer
          open
          variant="persistent"
          PaperProps={{
            sx: { width: DRAWER_WIDTH, bgcolor: "#1E2F40" },
          }}
        >
          {renderContent}
        </Drawer>
      </Hidden>
    </RootStyle>
  )
}
