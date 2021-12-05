import 'package:chegirma/data/model/categories_response.dart';
import 'package:chegirma/data/model/product_detail_response.dart';
import 'package:chegirma/data/model/product_detail_response.dart';
import 'package:chegirma/data/model/products_response.dart';
import 'package:chegirma/data/model/products_response.dart';
import 'package:chegirma/data/provider/api_client.dart';
import 'package:chegirma/data/provider/response_handler.dart';
import 'package:chegirma/data/provider/server_error.dart';
import 'package:dio/dio.dart';

class MainRepository {
  final apiClient = ApiClient();
  Future<ResponseHandler<CategoriesResponse>> getCategories() async {
    CategoriesResponse response;
    try {
      response = await apiClient.getCategories();
    } catch (error, stacktrace) {
      print("Exception occurred: $error stacktrace: $stacktrace");
      return ResponseHandler()..setException(ServerError.withError(error: error as DioError));
    }
    return ResponseHandler()..data = response;
  }

  Future<ResponseHandler<ProductsResponse>> getProducts({required String categoryId}) async {
    ProductsResponse response;
    try {
      response = await apiClient.getProducts(categoryId);
    } catch (error, stacktrace) {
      print("Exception occurred: $error stacktrace: $stacktrace");
      return ResponseHandler()..setException(ServerError.withError(error: error as DioError));
    }
    return ResponseHandler()..data = response;
  }

  Future<ResponseHandler<ProductDetailResponse>> getProductDetail({required String productId}) async {
    ProductDetailResponse response;
    try {
      response = await apiClient.getProductDetail(productId);
    } catch (error, stacktrace) {
      print("Exception occurred: $error stacktrace: $stacktrace");
      return ResponseHandler()..setException(ServerError.withError(error: error as DioError));
    }
    return ResponseHandler()..data = response;
  }

}
