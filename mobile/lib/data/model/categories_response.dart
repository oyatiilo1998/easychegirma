class CategoriesResponse {
  CategoriesResponse({
      this.count, 
      this.categories,});

  CategoriesResponse.fromJson(dynamic json) {
    count = json['count'];
    if (json['categories'] != null) {
      categories = [];
      json['categories'].forEach((v) {
        categories?.add(Categories.fromJson(v));
      });
    }
  }
  int? count;
  List<Categories>? categories;

  Map<String, dynamic> toJson() {
    final map = <String, dynamic>{};
    map['count'] = count;
    if (categories != null) {
      map['categories'] = categories?.map((v) => v.toJson()).toList();
    }
    return map;
  }

}

class Categories {
  Categories({
      this.id, 
      this.ruName, 
      this.name, 
      this.code,});

  Categories.fromJson(dynamic json) {
    id = json['id'];
    ruName = json['ru_name'];
    name = json['name'];
    code = json['code'];
  }
  String? id;
  String? ruName;
  String? name;
  int? code;

  Map<String, dynamic> toJson() {
    final map = <String, dynamic>{};
    map['id'] = id;
    map['ru_name'] = ruName;
    map['name'] = name;
    map['code'] = code;
    return map;
  }

}