// Copyright (C) 2018 go-nebulas authors
//
// This file is part of the go-nebulas library.
//
// the go-nebulas library is free software: you can redistribute it and/or
// modify
// it under the terms of the GNU General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// the go-nebulas library is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
// GNU General Public License for more details.
//
// You should have received a copy of the GNU General Public License
// along with the go-nebulas library.  If not, see
// <http://www.gnu.org/licenses/>.
//

#include "common/configuration.h"
#include "fs/trie/trie.h"
#include <gtest/gtest.h>

TEST(test_fs, key_to_route) {

  neb::util::bytes key(std::initializer_list<neb::byte_t>({0xa1, 0xf2}));
  neb::util::bytes route_expect(
      std::initializer_list<neb::byte_t>({0xa, 0x1, 0xf, 0x2}));

  auto route_actual = neb::fs::trie::key_to_route(key);
  ASSERT_EQ(route_expect.size(), route_actual.size());
  for (size_t i = 0; i < route_actual.size(); i++) {
    ASSERT_EQ(route_expect[i], route_actual[i]);
  }
}

TEST(test_fs, route_to_key) {

  neb::util::bytes route(
      std::initializer_list<neb::byte_t>({0xa, 0x1, 0xf, 0x2}));
  neb::util::bytes key_expect(std::initializer_list<neb::byte_t>({0xa1, 0xf2}));

  auto key_actual = neb::fs::trie::route_to_key(route);
  ASSERT_EQ(key_expect.size(), key_actual.size());
  for (size_t i = 0; i < key_actual.size(); i++) {
    ASSERT_EQ(key_expect[i], key_actual[i]);
  }
}

TEST(test_fs, get_trie_node) {
  std::string neb_db_path = neb::configuration::instance().neb_db_dir();
  neb::fs::rocksdb_storage rs;
  rs.open_database(neb_db_path, neb::fs::storage_open_for_readonly);
  neb::fs::trie t(&rs);

  std::string root_hash_str =
      "fd398c1351e60c44f44038fef233525f59d4fee3a4b4cf5546653a3a188b04db";
  rs.close_database();
}
