/*
 * Licensed to the Apache Software Foundation (ASF) under one or more
 * contributor license agreements.  See the NOTICE file distributed with
 * this work for additional information regarding copyright ownership.
 * The ASF licenses this file to You under the Apache License, Version 2.0
 * (the "License"); you may not use this file except in compliance with
 * the License.  You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package mysql

import (
	"strings"
)

var MySQLKeyword = map[string]string{
	"ACCESSIBLE":                    "ACCESSIBLE",
	"ADD":                           "ADD",
	"ALL":                           "ALL",
	"ALTER":                         "ALTER",
	"ANALYZE":                       "ANALYZE",
	"AND":                           "AND",
	"ARRAY":                         "ARRAY",
	"AS":                            "AS",
	"ASC":                           "ASC",
	"ASENSITIVE":                    "ASENSITIVE",
	"BEFORE":                        "BEFORE",
	"BETWEEN":                       "BETWEEN",
	"BIGINT":                        "BIGINT",
	"BINARY":                        "BINARY",
	"BLOB":                          "BLOB",
	"BOTH":                          "BOTH",
	"BY":                            "BY",
	"CALL":                          "CALL",
	"CASCADE":                       "CASCADE",
	"CASE":                          "CASE",
	"CHANGE":                        "CHANGE",
	"CHAR":                          "CHAR",
	"CHARACTER":                     "CHARACTER",
	"CHECK":                         "CHECK",
	"COLLATE":                       "COLLATE",
	"COLUMN":                        "COLUMN",
	"CONDITION":                     "CONDITION",
	"CONSTRAINT":                    "CONSTRAINT",
	"CONTINUE":                      "CONTINUE",
	"CONVERT":                       "CONVERT",
	"CREATE":                        "CREATE",
	"CROSS":                         "CROSS",
	"CUBE":                          "CUBE",
	"CUME_DIST":                     "CUME_DIST",
	"CURRENT_DATE":                  "CURRENT_DATE",
	"CURRENT_TIME":                  "CURRENT_TIME",
	"CURRENT_TIMESTAMP":             "CURRENT_TIMESTAMP",
	"CURRENT_USER":                  "CURRENT_USER",
	"CURSOR":                        "CURSOR",
	"DATABASE":                      "DATABASE",
	"DATABASES":                     "DATABASES",
	"DAY_HOUR":                      "DAY_HOUR",
	"DAY_MICROSECOND":               "DAY_MICROSECOND",
	"DAY_MINUTE":                    "DAY_MINUTE",
	"DAY_SECOND":                    "DAY_SECOND",
	"DEC":                           "DEC",
	"DECIMAL":                       "DECIMAL",
	"DECLARE":                       "DECLARE",
	"DEFAULT":                       "DEFAULT",
	"DELAYED":                       "DELAYED",
	"DELETE":                        "DELETE",
	"DENSE_RANK":                    "DENSE_RANK",
	"DESC":                          "DESC",
	"DESCRIBE":                      "DESCRIBE",
	"DETERMINISTIC":                 "DETERMINISTIC",
	"DISTINCT":                      "DISTINCT",
	"DISTINCTROW":                   "DISTINCTROW",
	"DIV":                           "DIV",
	"DOUBLE":                        "DOUBLE",
	"DROP":                          "DROP",
	"DUAL":                          "DUAL",
	"EACH":                          "EACH",
	"ELSE":                          "ELSE",
	"ELSEIF":                        "ELSEIF",
	"EMPTY":                         "EMPTY",
	"ENCLOSED":                      "ENCLOSED",
	"ESCAPED":                       "ESCAPED",
	"EXCEPT":                        "EXCEPT",
	"EXISTS":                        "EXISTS",
	"EXIT":                          "EXIT",
	"EXPLAIN":                       "EXPLAIN",
	"FALSE":                         "FALSE",
	"FETCH":                         "FETCH",
	"FIRST_VALUE":                   "FIRST_VALUE",
	"FLOAT":                         "FLOAT",
	"FLOAT4":                        "FLOAT4",
	"FLOAT8":                        "FLOAT8",
	"FOR":                           "FOR",
	"FORCE":                         "FORCE",
	"FOREIGN":                       "FOREIGN",
	"FROM":                          "FROM",
	"FULLTEXT":                      "FULLTEXT",
	"FUNCTION":                      "FUNCTION",
	"GENERATED":                     "GENERATED",
	"GET":                           "GET",
	"GRANT":                         "GRANT",
	"GROUP":                         "GROUP",
	"GROUPING":                      "GROUPING",
	"GROUPS":                        "GROUPS",
	"HAVING":                        "HAVING",
	"HIGH_PRIORITY":                 "HIGH_PRIORITY",
	"HOUR_MICROSECOND":              "HOUR_MICROSECOND",
	"HOUR_MINUTE":                   "HOUR_MINUTE",
	"HOUR_SECOND":                   "HOUR_SECOND",
	"IF":                            "IF",
	"IGNORE":                        "IGNORE",
	"IN":                            "IN",
	"INDEX":                         "INDEX",
	"INFILE":                        "INFILE",
	"INNER":                         "INNER",
	"INOUT":                         "INOUT",
	"INSENSITIVE":                   "INSENSITIVE",
	"INSERT":                        "INSERT",
	"INT":                           "INT",
	"INT1":                          "INT1",
	"INT2":                          "INT2",
	"INT3":                          "INT3",
	"INT4":                          "INT4",
	"INT8":                          "INT8",
	"INTEGER":                       "INTEGER",
	"INTERVAL":                      "INTERVAL",
	"INTO":                          "INTO",
	"IO_AFTER_GTIDS":                "IO_AFTER_GTIDS",
	"IO_BEFORE_GTIDS":               "IO_BEFORE_GTIDS",
	"IS":                            "IS",
	"ITERATE":                       "ITERATE",
	"JOIN":                          "JOIN",
	"JSON_TABLE":                    "JSON_TABLE",
	"KEY":                           "KEY",
	"KEYS":                          "KEYS",
	"KILL":                          "KILL",
	"LAG":                           "LAG",
	"LAST_VALUE":                    "LAST_VALUE",
	"LATERAL":                       "LATERAL",
	"LEAD":                          "LEAD",
	"LEADING":                       "LEADING",
	"LEAVE":                         "LEAVE",
	"LEFT":                          "LEFT",
	"LIKE":                          "LIKE",
	"LIMIT":                         "LIMIT",
	"LINEAR":                        "LINEAR",
	"LINES":                         "LINES",
	"LOAD":                          "LOAD",
	"LOCALTIME":                     "LOCALTIME",
	"LOCALTIMESTAMP":                "LOCALTIMESTAMP",
	"LOCK":                          "LOCK",
	"LONG":                          "LONG",
	"LONGBLOB":                      "LONGBLOB",
	"LONGTEXT":                      "LONGTEXT",
	"LOOP":                          "LOOP",
	"LOW_PRIORITY":                  "LOW_PRIORITY",
	"MASTER_BIND":                   "MASTER_BIND",
	"MASTER_SSL_VERIFY_SERVER_CERT": "MASTER_SSL_VERIFY_SERVER_CERT",
	"MATCH":                         "MATCH",
	"MAXVALUE":                      "MAXVALUE",
	"MEDIUMBLOB":                    "MEDIUMBLOB",
	"MEDIUMINT":                     "MEDIUMINT",
	"MEDIUMTEXT":                    "MEDIUMTEXT",
	"MEMBER":                        "MEMBER",
	"MIDDLEINT":                     "MIDDLEINT",
	"MINUTE_MICROSECOND":            "MINUTE_MICROSECOND",
	"MINUTE_SECOND":                 "MINUTE_SECOND",
	"MOD":                           "MOD",
	"MODIFIES":                      "MODIFIES",
	"NATURAL":                       "NATURAL",
	"NOT":                           "NOT",
	"NO_WRITE_TO_BINLOG":            "NO_WRITE_TO_BINLOG",
	"NTH_VALUE":                     "NTH_VALUE",
	"NTILE":                         "NTILE",
	"NULL":                          "NULL",
	"NUMERIC":                       "NUMERIC",
	"OF":                            "OF",
	"ON":                            "ON",
	"OPTIMIZE":                      "OPTIMIZE",
	"OPTIMIZER_COSTS":               "OPTIMIZER_COSTS",
	"OPTION":                        "OPTION",
	"OPTIONALLY":                    "OPTIONALLY",
	"OR":                            "OR",
	"ORDER":                         "ORDER",
	"OUT":                           "OUT",
	"OUTER":                         "OUTER",
	"OUTFILE":                       "OUTFILE",
	"OVER":                          "OVER",
	"PARTITION":                     "PARTITION",
	"PERCENT_RANK":                  "PERCENT_RANK",
	"PRECISION":                     "PRECISION",
	"PRIMARY":                       "PRIMARY",
	"PROCEDURE":                     "PROCEDURE",
	"PURGE":                         "PURGE",
	"RANGE":                         "RANGE",
	"RANK":                          "RANK",
	"READ":                          "READ",
	"READS":                         "READS",
	"READ_WRITE":                    "READ_WRITE",
	"REAL":                          "REAL",
	"RECURSIVE":                     "RECURSIVE",
	"REFERENCES":                    "REFERENCES",
	"REGEXP":                        "REGEXP",
	"RELEASE":                       "RELEASE",
	"RENAME":                        "RENAME",
	"REPEAT":                        "REPEAT",
	"REPLACE":                       "REPLACE",
	"REQUIRE":                       "REQUIRE",
	"RESIGNAL":                      "RESIGNAL",
	"RESTRICT":                      "RESTRICT",
	"RETURN":                        "RETURN",
	"REVOKE":                        "REVOKE",
	"RIGHT":                         "RIGHT",
	"RLIKE":                         "RLIKE",
	"ROW":                           "ROW",
	"ROWS":                          "ROWS",
	"ROW_NUMBER":                    "ROW_NUMBER",
	"SCHEMA":                        "SCHEMA",
	"SCHEMAS":                       "SCHEMAS",
	"SECOND_MICROSECOND":            "SECOND_MICROSECOND",
	"SELECT":                        "SELECT",
	"SENSITIVE":                     "SENSITIVE",
	"SEPARATOR":                     "SEPARATOR",
	"SET":                           "SET",
	"SHOW":                          "SHOW",
	"SIGNAL":                        "SIGNAL",
	"SMALLINT":                      "SMALLINT",
	"SPATIAL":                       "SPATIAL",
	"SPECIFIC":                      "SPECIFIC",
	"SQL":                           "SQL",
	"SQLEXCEPTION":                  "SQLEXCEPTION",
	"SQLSTATE":                      "SQLSTATE",
	"SQLWARNING":                    "SQLWARNING",
	"SQL_BIG_RESULT":                "SQL_BIG_RESULT",
	"SQL_CALC_FOUND_ROWS":           "SQL_CALC_FOUND_ROWS",
	"SQL_SMALL_RESULT":              "SQL_SMALL_RESULT",
	"SSL":                           "SSL",
	"STARTING":                      "STARTING",
	"STORED":                        "STORED",
	"STRAIGHT_JOIN":                 "STRAIGHT_JOIN",
	"SYSTEM":                        "SYSTEM",
	"TABLE":                         "TABLE",
	"TERMINATED":                    "TERMINATED",
	"THEN":                          "THEN",
	"TINYBLOB":                      "TINYBLOB",
	"TINYINT":                       "TINYINT",
	"TINYTEXT":                      "TINYTEXT",
	"TO":                            "TO",
	"TRAILING":                      "TRAILING",
	"TRIGGER":                       "TRIGGER",
	"TRUE":                          "TRUE",
	"UNDO":                          "UNDO",
	"UNION":                         "UNION",
	"UNIQUE":                        "UNIQUE",
	"UNLOCK":                        "UNLOCK",
	"UNSIGNED":                      "UNSIGNED",
	"UPDATE":                        "UPDATE",
	"USAGE":                         "USAGE",
	"USE":                           "USE",
	"USING":                         "USING",
	"UTC_DATE":                      "UTC_DATE",
	"UTC_TIME":                      "UTC_TIME",
	"UTC_TIMESTAMP":                 "UTC_TIMESTAMP",
	"VALUES":                        "VALUES",
	"VARBINARY":                     "VARBINARY",
	"VARCHAR":                       "VARCHAR",
	"VARCHARACTER":                  "VARCHARACTER",
	"VARYING":                       "VARYING",
	"VIRTUAL":                       "VIRTUAL",
	"WHEN":                          "WHEN",
	"WHERE":                         "WHERE",
	"WHILE":                         "WHILE",
	"WINDOW":                        "WINDOW",
	"WITH":                          "WITH",
	"WRITE":                         "WRITE",
	"XOR":                           "XOR",
	"YEAR_MONTH":                    "YEAR_MONTH",
	"ZEROFILL":                      "ZEROFILL",
}

func Check(fieldOrTableName string) bool {
	_, ok := MySQLKeyword[fieldOrTableName]
	if ok {
		return true
	}
	if fieldOrTableName != "" {
		fieldOrTableName = strings.ToUpper(fieldOrTableName)
	}
	_, ok = MySQLKeyword[fieldOrTableName]
	return ok
}

func CheckEscape(fieldOrTableName string) bool {
	return Check(fieldOrTableName)
}

func CheckAndReplace(fieldOrTableName string) string {
	if Check(fieldOrTableName) {
		return "`" + fieldOrTableName + "`"
	} else {
		return fieldOrTableName
	}
}
