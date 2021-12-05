import 'dart:async';
import 'dart:io';
import 'package:chegirma/data/model/categories_response.dart';
import 'package:chegirma/data/model/product_detail_response.dart';
import 'package:chegirma/data/model/products_response.dart';
import 'package:dio/adapter.dart';
import 'package:dio/dio.dart';
import 'package:retrofit/http.dart';
import 'package:retrofit/retrofit.dart';

part 'api_client.g.dart';

@RestApi(baseUrl: 'https://api.chegirma.uz/v1/')
abstract class ApiClient {
  factory ApiClient({String baseUrl = 'https://api.chegirma.uz/v1/'}) {
    Dio dio = Dio(BaseOptions(followRedirects: false));
    (dio.httpClientAdapter as DefaultHttpClientAdapter).onHttpClientCreate = (HttpClient client) {
      client.badCertificateCallback = (X509Certificate cert, String host, int port) => true;
      return client;
    };
    dio.interceptors.add(LogInterceptor(responseBody: true, requestBody: true, requestHeader: true, request: true, error: true));
    dio.options = BaseOptions(
      receiveTimeout: 60000,
      connectTimeout: 60000,
      sendTimeout: 60000,
    );
    return _ApiClient(dio, baseUrl: baseUrl);
  }

  @GET('category')
  Future<CategoriesResponse> getCategories();

  @GET('product')
  Future<ProductsResponse> getProducts(
    @Query("category_id") String categoryId,
  );

  @GET('product/{product_id}')
  Future<ProductDetailResponse> getProductDetail(
    @Path("product_id") String productId,
  );
}
