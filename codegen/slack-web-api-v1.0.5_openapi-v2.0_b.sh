perl -p -i -e 's/"\$ref":\s*"#\/definitions\/defs_ok_true"/$1"type": "boolean"/g' swagger_spec.json
perl -p -i -e 's/"\$ref":\s*"#\/definitions\/defs_ok_false"/$1"type": "boolean"/g' swagger_spec.json
perl -p -i -e 's/"\$ref":\s*"#\/definitions\/defs_invite_id"/$1"type": "integer",\n"format":"int64"/g' swagger_spec.json
perl -p -i -e 's/"\$ref":\s*"#\/definitions\/defs_[^"]+"/$1"type": "string"/g' swagger_spec.json