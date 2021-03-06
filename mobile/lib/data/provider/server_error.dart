import 'package:dio/dio.dart' hide Headers;

class ServerError implements Exception {
  int? _errorCode;
  String _errorMessage = "";

  ServerError.withError({DioError? error}) {
    _handleError(error);
  }

  int getErrorCode() => _errorCode??500;

  String getErrorMessage() => _errorMessage;

  _handleError(DioError? error) {
    _errorCode = error!.response?.statusCode ?? 500;
    switch (error.type) {
      case DioErrorType.connectTimeout:
        _errorMessage = "Connection timeout";
        break;
      case DioErrorType.sendTimeout:
        _errorMessage = "Connection timeout";
        break;
      case DioErrorType.receiveTimeout:
        _errorMessage = "Connection timeout";
        break;
      case DioErrorType.response:
        {
          if (error.response?.data['error'] is Map<String, dynamic>) {
            _errorMessage = error.response!.data['error']['message'].toString();
          } else {
            _errorMessage = error.response!.data['message'].toString();
          }
          break;
        }
      case DioErrorType.cancel:
        _errorMessage = "Canceled";
        break;
      case DioErrorType.other:
        _errorMessage = "Something wrong";
        break;
    }
    return _errorMessage;
  }
}
