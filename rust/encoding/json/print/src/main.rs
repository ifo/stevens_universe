#[macro_use]
extern crate serde_derive;

extern crate serde;
extern crate serde_json;

use serde_json::Value;

#[derive(Serialize, Deserialize, Debug)]
struct Obj {
    one: u32,
    two: Vec<u32>,
}

fn main() {
    let input = r#"{"one":1,"two":[1,2]}"#;

    let v: Value = serde_json::from_str(input)
        .expect("Couldn't parse it.");

    println!("value                : {}", v);

    let obj: Obj = serde_json::from_str(&input).unwrap();

    println!("object               : {:?}", obj);

    let output = serde_json::json!(obj);

    println!("string with json!    : {}", output);

    let to_string = serde_json::to_string(&obj).unwrap();

    println!("string with to_string: {}", to_string);

    let pretty = serde_json::to_string_pretty(&obj).unwrap();

    println!("pretty printed       :\n{}", pretty);
}
