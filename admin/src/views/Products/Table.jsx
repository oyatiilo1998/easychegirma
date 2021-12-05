import {
  experimentalStyled,
  Pagination,
  Paper,
  Table,
  TableBody,
  TableCell,
  TableContainer,
  TableHead,
  TableRow,
} from "@mui/material"
import { useEffect, useState } from "react"
import { useSelector } from "react-redux"
import { useHistory } from "react-router"
import NoDataComponent from "../../components/NoDataComponent"
import TableLoader from "../../components/TableLoader"
import categoryService from "../../services/categoryService"
import productService from "../../services/productService"

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

const ProductsTable = () => {
  const history = useHistory()

  const [tableData, setTableData] = useState(null)
  const [loader, setLoader] = useState(true)
  const [pageCount, setPageCount] = useState(1)
  const [currentPage, setCurrentPage] = useState(1)
  const userId = useSelector(state => state.auth.userId)

  const fetchTableData = () => {
    setLoader(true)
    productService
      .getProductList({
        limit: 10,
        page: currentPage,
        owner_id: userId
      })
      .then((res) => {
        setTableData(res?.products)
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
              <StledTableCell>Скидка</StledTableCell>
              <StledTableCell>Цена до скидки</StledTableCell>
              <StledTableCell>Цена после скидки</StledTableCell>
            </TableRow>
          </TableHead>
          <TableBody style={{ position: "relative" }}>
            {tableData?.map((data, index) => (
              <StyledTableRow
                onClick={() => history.push(`/main/products/edit/${data.id}`)}
              >
                <StledTableCell style={{ width: "50px" }}>
                  {index + 1}
                </StledTableCell>
                <StledTableCell>{data?.name}</StledTableCell>
                <StledTableCell>{data?.discount_amount} %</StledTableCell>
                <StledTableCell>{data?.price_before} сум</StledTableCell>
                <StledTableCell>{data?.price_after} сум</StledTableCell>
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

export default ProductsTable
