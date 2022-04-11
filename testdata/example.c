//go:build ignore

#include <stdio.h>
#include <yyjson.h>

int
main() {
  yyjson_read_err err;
  yyjson_doc *doc = yyjson_read_file("testdata/example.json",
      YYJSON_READ_NOFLAG,
      NULL,
      &err);
  if (!doc) {
    fprintf(stderr, "%s\n", err.msg);
    exit(1);
  }

  yyjson_val *root = yyjson_doc_get_root(doc);

  yyjson_val *title = yyjson_obj_get(root, "title");
  printf("title: %s\n", yyjson_get_str(title));

  yyjson_val *entries = yyjson_obj_get(root, "entries");

  size_t idx, max;
  yyjson_val *entry;
  yyjson_arr_foreach(entries, idx, max, entry) {
    title = yyjson_obj_get(entry, "title");
    printf("%s\n", yyjson_get_str(title));
  }

  yyjson_doc_free(doc); 

  return 0;
}
