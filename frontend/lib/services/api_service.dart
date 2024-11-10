import 'package:dio/dio.dart';
import 'package:myshow/models/user.dart';
import 'package:myshow/models/event.dart';
import 'package:shared_preferences/shared_preferences.dart';

class ApiService {
  final Dio _dio = Dio(BaseOptions(baseUrl: 'http://192.168.1.15:8080/api'));
  String? _jwtToken;

  Future<void> _setAuthHeader() async {
    if (_jwtToken == null) {
      final prefs = await SharedPreferences.getInstance();
      _jwtToken = prefs.getString('jwtToken');
    }
    if (_jwtToken != null) {
      _dio.options.headers['Authorization'] = 'Bearer $_jwtToken';
    }
  }

  Future<User> loginUser(String username, String password) async {
    final response = await _dio.post('/users/login',
        data: {'username': username, 'password': password});
    _jwtToken = response.data['token'];
    final prefs = await SharedPreferences.getInstance();
    await prefs.setString('jwtToken', _jwtToken!);
    return User.fromJson(response.data['user']);
  }

  Future<User> registerUser(
      String username, String email, String password) async {
    final response = await _dio.post('/users/register',
        data: {'username': username, 'email': email, 'password': password});
    return User.fromJson(response.data['user']);
  }

  Future<List<Event>> getEvents() async {
    await _setAuthHeader();
    final response = await _dio.get('/events');
    return (response.data as List)
        .map((event) => Event.fromJson(event))
        .toList();
  }

  Future<Event> getEvent(int id) async {
    await _setAuthHeader();
    final response = await _dio.get('/events/$id');
    return Event.fromJson(response.data);
  }

  Future<void> deleteEvent(int id) async {
    await _setAuthHeader();
    await _dio.delete('/events/$id');
  }
}
