1. Sort all parameter by keys in alphabetical order.
2. Append them (if the value is not null or empty) in key1=value1&key2=value2.
3. Append secret key, e.g.: key1=value1&key2=value2SECRET.
4. Calculate the hash by using MD5.
