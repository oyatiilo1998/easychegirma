import LocalGroceryStoreIcon from '@mui/icons-material/LocalGroceryStore';
import TocIcon from "@mui/icons-material/Toc"
import AssignmentIndIcon from '@mui/icons-material/AssignmentInd';
// ----------------------------------------------------------------------

const sidebarConfig = [
  {
    subheader: "",
    items: [
      {
        title: "Категории",
        href: "/main/categories",
        icon: <TocIcon />,
      },
      {
        title: "Продукты",
        href: "/main/products",
        icon: <LocalGroceryStoreIcon />,
      },
      {
        title: "Пользователи",
        href: "/main/users",
        icon: <AssignmentIndIcon />,
      },
    ],
  },
]

export default sidebarConfig
