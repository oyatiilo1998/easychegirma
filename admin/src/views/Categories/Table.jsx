import {
  experimentalStyled,
  Pagination,
  Paper,
  Switch,
  Table,
  TableBody,
  TableCell,
  TableContainer,
  TableHead,
  TableRow,
} from "@mui/material"
import { useEffect, useState } from "react"
import { useHistory } from "react-router"
import NoDataComponent from "../../components/NoDataComponent"
import TableLoader from "../../components/TableLoader"
import categoryService from "../../services/categoryService"

const StledTableCell = experimentalStyled(TableCell)(({ theme }) => ({
  textAlign: "center",
}))

const StyledTableRow = experimentalStyled(TableRow)(({ theme }) => ({
  "&:nth-of-type(odd)": {
    // backgroundColor: theme.palette.action.hover,
  },
  "&:hover": {
    backgroundColor: theme.palette.action.hover,
    cursor: "pointer",
  },
  // hide last border
  "&:last-child td, &:last-child th": {
    border: 0,
  },
}))

const StyledTableContainer = experimentalStyled(TableContainer)(
  ({ theme }) => ({
    height: "calc(100vh - 200px)",
    position: "relative",
  })
)

const CategoriesTable = () => {
  const history = useHistory()

  const [tableData, setTableData] = useState(null)
  const [loader, setLoader] = useState(true)
  const [pageCount, setPageCount] = useState(1)
  const [currentPage, setCurrentPage] = useState(1)

  const fetchTableData = () => {
    setLoader(true)
    categoryService
      .getCategoryList({
        limit: 10,
        page: currentPage
      })
      .then((res) => {
        setTableData(res?.categories)
        setPageCount(Math.ceil(res?.count / 10))
      })
      .finally(() => setLoader(false))
  }

  useEffect(() => {
    fetchTableData()
  }, [currentPage])

  return (
    <Paper elevation={12} style={{ padding: "20px" }}>
      <StyledTableContainer>
        <Table stickyHeader>
          <TableHead>
            <TableRow>
              <StledTableCell>No</StledTableCell>
              <StledTableCell>Название</StledTableCell>
            </TableRow>
          </TableHead>
          <TableBody style={{ position: "relative" }}>
            {tableData?.map((data, index) => (
              <StyledTableRow
                onClick={() => history.push(`/main/categories/edit/${data.id}`)}
              >
                <StledTableCell style={{ width: "50px" }}>
                  {index + 1}
                </StledTableCell>
                <StledTableCell>{data?.name}</StledTableCell>
              </StyledTableRow>
            ))}

            <TableLoader columnsCount={2} isVisible={loader} />
            <NoDataComponent isVisible={!loader && !tableData?.length} />
          </TableBody>
        </Table>
      </StyledTableContainer>

      <div className="pagination-container">
        <Pagination
          count={pageCount}
          color="primary"
          page={currentPage}
          onChange={(e, val) => setCurrentPage(val)}
        />
      </div>
    </Paper>
  )
}

export default CategoriesTable
