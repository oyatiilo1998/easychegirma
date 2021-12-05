class ProductsResponse {
  ProductsResponse({
      this.count, 
      this.products,});

  ProductsResponse.fromJson(dynamic json) {
    count = json['count'];
    if (json['products'] != null) {
      products = [];
      json['products'].forEach((v) {
        products?.add(Products.fromJson(v));
      });
    }
  }
  int? count;
  List<Products>? products;

  Map<String, dynamic> toJson() {
    final map = <String, dynamic>{};
    map['count'] = count;
    if (products != null) {
      map['products'] = products?.map((v) => v.toJson()).toList();
    }
    return map;
  }

}

class Products {
  Products({
      this.id, 
      this.url, 
      this.name, 
      this.imageUrl, 
      this.discountAmount, 
      this.priceBefore, 
      this.priceAfter, 
      this.category, 
      this.fromTime, 
      this.toTime, 
      this.owner,});

  Products.fromJson(dynamic json) {
    id = json['id'];
    url = json['url'];
    name = json['name'];
    imageUrl = json['image_url'];
    discountAmount = json['discount_amount'];
    priceBefore = json['price_before'];
    priceAfter = json['price_after'];
    category = json['category'] != null ? Category.fromJson(json['category']) : null;
    fromTime = json['from_time'];
    toTime = json['to_time'];
    owner = json['owner'] != null ? Owner.fromJson(json['owner']) : null;
  }
  String? id;
  String? url;
  String? name;
  String? imageUrl;
  int? discountAmount;
  int? priceBefore;
  int? priceAfter;
  Category? category;
  String? fromTime;
  String? toTime;
  Owner? owner;

  Map<String, dynamic> toJson() {
    final map = <String, dynamic>{};
    map['id'] = id;
    map['url'] = url;
    map['name'] = name;
    map['image_url'] = imageUrl;
    map['discount_amount'] = discountAmount;
    map['price_before'] = priceBefore;
    map['price_after'] = priceAfter;
    if (category != null) {
      map['category'] = category?.toJson();
    }
    map['from_time'] = fromTime;
    map['to_time'] = toTime;
    if (owner != null) {
      map['owner'] = owner?.toJson();
    }
    return map;
  }

}

class Owner {
  Owner({
      this.id, 
      this.name, 
      this.password, 
      this.login, 
      this.imageUrl, 
      this.link,});

  Owner.fromJson(dynamic json) {
    id = json['id'];
    name = json['name'];
    password = json['password'];
    login = json['login'];
    imageUrl = json['image_url'];
    link = json['link'];
  }
  String? id;
  String? name;
  String? password;
  String? login;
  String? imageUrl;
  String? link;

  Map<String, dynamic> toJson() {
    final map = <String, dynamic>{};
    map['id'] = id;
    map['name'] = name;
    map['password'] = password;
    map['login'] = login;
    map['image_url'] = imageUrl;
    map['link'] = link;
    return map;
  }

}

class Category {
  Category({
      this.id, 
      this.ruName, 
      this.name, 
      this.code,});

  Category.fromJson(dynamic json) {
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