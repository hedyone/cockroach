// Code generated by all_keywords.awk. DO NOT EDIT.
// GENERATED FILE DO NOT EDIT

package parser

var keywords = map[string]struct {
	tok int
	cat string
}{
	"ACTION":                    {ACTION, "U"},
	"ADD":                       {ADD, "U"},
	"ALL":                       {ALL, "R"},
	"ALTER":                     {ALTER, "U"},
	"ANALYSE":                   {ANALYSE, "R"},
	"ANALYZE":                   {ANALYZE, "R"},
	"AND":                       {AND, "R"},
	"ANNOTATE_TYPE":             {ANNOTATE_TYPE, "C"},
	"ANY":                       {ANY, "R"},
	"ARRAY":                     {ARRAY, "R"},
	"AS":                        {AS, "R"},
	"ASC":                       {ASC, "R"},
	"ASYMMETRIC":                {ASYMMETRIC, "R"},
	"AT":                        {AT, "U"},
	"BACKUP":                    {BACKUP, "U"},
	"BEGIN":                     {BEGIN, "U"},
	"BETWEEN":                   {BETWEEN, "C"},
	"BIGINT":                    {BIGINT, "C"},
	"BIGSERIAL":                 {BIGSERIAL, "C"},
	"BIT":                       {BIT, "C"},
	"BLOB":                      {BLOB, "U"},
	"BOOL":                      {BOOL, "C"},
	"BOOLEAN":                   {BOOLEAN, "C"},
	"BOTH":                      {BOTH, "R"},
	"BY":                        {BY, "U"},
	"BYTEA":                     {BYTEA, "C"},
	"BYTES":                     {BYTES, "C"},
	"CANCEL":                    {CANCEL, "U"},
	"CASCADE":                   {CASCADE, "U"},
	"CASE":                      {CASE, "R"},
	"CAST":                      {CAST, "R"},
	"CHAR":                      {CHAR, "C"},
	"CHARACTER":                 {CHARACTER, "C"},
	"CHARACTERISTICS":           {CHARACTERISTICS, "C"},
	"CHECK":                     {CHECK, "R"},
	"CLUSTER":                   {CLUSTER, "U"},
	"COALESCE":                  {COALESCE, "C"},
	"COLLATE":                   {COLLATE, "R"},
	"COLLATION":                 {COLLATION, "T"},
	"COLUMN":                    {COLUMN, "R"},
	"COLUMNS":                   {COLUMNS, "U"},
	"COMMIT":                    {COMMIT, "U"},
	"COMMITTED":                 {COMMITTED, "U"},
	"CONFLICT":                  {CONFLICT, "U"},
	"CONSTRAINT":                {CONSTRAINT, "R"},
	"CONSTRAINTS":               {CONSTRAINTS, "U"},
	"COPY":                      {COPY, "U"},
	"COVERING":                  {COVERING, "U"},
	"CREATE":                    {CREATE, "R"},
	"CROSS":                     {CROSS, "T"},
	"CSV":                       {CSV, "U"},
	"CUBE":                      {CUBE, "U"},
	"CURRENT":                   {CURRENT, "U"},
	"CURRENT_CATALOG":           {CURRENT_CATALOG, "R"},
	"CURRENT_DATE":              {CURRENT_DATE, "R"},
	"CURRENT_ROLE":              {CURRENT_ROLE, "R"},
	"CURRENT_SCHEMA":            {CURRENT_SCHEMA, "R"},
	"CURRENT_TIME":              {CURRENT_TIME, "R"},
	"CURRENT_TIMESTAMP":         {CURRENT_TIMESTAMP, "R"},
	"CURRENT_USER":              {CURRENT_USER, "R"},
	"CYCLE":                     {CYCLE, "U"},
	"DATA":                      {DATA, "U"},
	"DATABASE":                  {DATABASE, "U"},
	"DATABASES":                 {DATABASES, "U"},
	"DATE":                      {DATE, "C"},
	"DAY":                       {DAY, "U"},
	"DEALLOCATE":                {DEALLOCATE, "U"},
	"DEC":                       {DEC, "C"},
	"DECIMAL":                   {DECIMAL, "C"},
	"DEFAULT":                   {DEFAULT, "R"},
	"DEFERRABLE":                {DEFERRABLE, "R"},
	"DELETE":                    {DELETE, "U"},
	"DESC":                      {DESC, "R"},
	"DISCARD":                   {DISCARD, "U"},
	"DISTINCT":                  {DISTINCT, "R"},
	"DO":                        {DO, "R"},
	"DOUBLE":                    {DOUBLE, "U"},
	"DROP":                      {DROP, "U"},
	"ELSE":                      {ELSE, "R"},
	"ENCODING":                  {ENCODING, "U"},
	"END":                       {END, "R"},
	"EXCEPT":                    {EXCEPT, "R"},
	"EXECUTE":                   {EXECUTE, "U"},
	"EXISTS":                    {EXISTS, "C"},
	"EXPERIMENTAL_FINGERPRINTS": {EXPERIMENTAL_FINGERPRINTS, "U"},
	"EXPLAIN":                   {EXPLAIN, "U"},
	"EXTRACT":                   {EXTRACT, "C"},
	"EXTRACT_DURATION":          {EXTRACT_DURATION, "C"},
	"FALSE":                     {FALSE, "R"},
	"FAMILY":                    {FAMILY, "T"},
	"FETCH":                     {FETCH, "R"},
	"FILTER":                    {FILTER, "U"},
	"FIRST":                     {FIRST, "U"},
	"FLOAT":                     {FLOAT, "C"},
	"FLOAT4":                    {FLOAT4, "C"},
	"FLOAT8":                    {FLOAT8, "C"},
	"FOLLOWING":                 {FOLLOWING, "U"},
	"FOR":                       {FOR, "R"},
	"FORCE_INDEX":               {FORCE_INDEX, "U"},
	"FOREIGN":                   {FOREIGN, "R"},
	"FROM":                      {FROM, "R"},
	"FULL":                      {FULL, "T"},
	"GRANT":                     {GRANT, "R"},
	"GRANTS":                    {GRANTS, "U"},
	"GREATEST":                  {GREATEST, "C"},
	"GROUP":                     {GROUP, "R"},
	"GROUPING":                  {GROUPING, "C"},
	"HAVING":                    {HAVING, "R"},
	"HIGH":                      {HIGH, "U"},
	"HOUR":                      {HOUR, "U"},
	"IF":                        {IF, "C"},
	"IFNULL":                    {IFNULL, "C"},
	"ILIKE":                     {ILIKE, "T"},
	"IMPORT":                    {IMPORT, "U"},
	"IN":                        {IN, "R"},
	"INCREMENTAL":               {INCREMENTAL, "U"},
	"INDEX":                     {INDEX, "R"},
	"INDEXES":                   {INDEXES, "U"},
	"INET":                      {INET, "C"},
	"INITIALLY":                 {INITIALLY, "R"},
	"INNER":                     {INNER, "T"},
	"INSERT":                    {INSERT, "U"},
	"INT":                       {INT, "C"},
	"INT2":                      {INT2, "C"},
	"INT2VECTOR":                {INT2VECTOR, "U"},
	"INT4":                      {INT4, "C"},
	"INT64":                     {INT64, "C"},
	"INT8":                      {INT8, "C"},
	"INTEGER":                   {INTEGER, "C"},
	"INTERLEAVE":                {INTERLEAVE, "U"},
	"INTERSECT":                 {INTERSECT, "R"},
	"INTERVAL":                  {INTERVAL, "C"},
	"INTO":                      {INTO, "R"},
	"IS":                        {IS, "T"},
	"ISOLATION":                 {ISOLATION, "U"},
	"JOB":                       {JOB, "U"},
	"JOBS":                      {JOBS, "U"},
	"JOIN":                      {JOIN, "T"},
	"KEY":                       {KEY, "U"},
	"KEYS":                      {KEYS, "U"},
	"KV":                        {KV, "U"},
	"LATERAL":                   {LATERAL, "R"},
	"LC_COLLATE":                {LC_COLLATE, "U"},
	"LC_CTYPE":                  {LC_CTYPE, "U"},
	"LEADING":                   {LEADING, "R"},
	"LEAST":                     {LEAST, "C"},
	"LEFT":                      {LEFT, "T"},
	"LEVEL":                     {LEVEL, "U"},
	"LIKE":                      {LIKE, "T"},
	"LIMIT":                     {LIMIT, "R"},
	"LOCAL":                     {LOCAL, "U"},
	"LOCALTIME":                 {LOCALTIME, "R"},
	"LOCALTIMESTAMP":            {LOCALTIMESTAMP, "R"},
	"LOW":                       {LOW, "U"},
	"MATCH":                     {MATCH, "U"},
	"MINUTE":                    {MINUTE, "U"},
	"MONTH":                     {MONTH, "U"},
	"NAME":                      {NAME, "C"},
	"NAMES":                     {NAMES, "U"},
	"NAN":                       {NAN, "U"},
	"NATURAL":                   {NATURAL, "T"},
	"NEXT":                      {NEXT, "U"},
	"NO":                        {NO, "U"},
	"NORMAL":                    {NORMAL, "U"},
	"NOT":                       {NOT, "R"},
	"NOTHING":                   {NOTHING, "R"},
	"NO_INDEX_JOIN":             {NO_INDEX_JOIN, "U"},
	"NULL":                      {NULL, "R"},
	"NULLIF":                    {NULLIF, "C"},
	"NULLS":                     {NULLS, "U"},
	"NUMERIC":                   {NUMERIC, "C"},
	"OF":                        {OF, "U"},
	"OFF":                       {OFF, "U"},
	"OFFSET":                    {OFFSET, "R"},
	"OID":                       {OID, "U"},
	"ON":                        {ON, "R"},
	"ONLY":                      {ONLY, "R"},
	"OPTIONS":                   {OPTIONS, "U"},
	"OR":                        {OR, "R"},
	"ORDER":                     {ORDER, "R"},
	"ORDINALITY":                {ORDINALITY, "U"},
	"OUT":                       {OUT, "C"},
	"OUTER":                     {OUTER, "T"},
	"OVER":                      {OVER, "U"},
	"OVERLAPS":                  {OVERLAPS, "T"},
	"OVERLAY":                   {OVERLAY, "C"},
	"PARENT":                    {PARENT, "U"},
	"PARTIAL":                   {PARTIAL, "U"},
	"PARTITION":                 {PARTITION, "U"},
	"PASSWORD":                  {PASSWORD, "U"},
	"PAUSE":                     {PAUSE, "U"},
	"PLACING":                   {PLACING, "R"},
	"PLANS":                     {PLANS, "U"},
	"POSITION":                  {POSITION, "C"},
	"PRECEDING":                 {PRECEDING, "U"},
	"PRECISION":                 {PRECISION, "C"},
	"PREPARE":                   {PREPARE, "U"},
	"PRIMARY":                   {PRIMARY, "R"},
	"PRIORITY":                  {PRIORITY, "U"},
	"QUERIES":                   {QUERIES, "U"},
	"QUERY":                     {QUERY, "U"},
	"RANGE":                     {RANGE, "U"},
	"READ":                      {READ, "U"},
	"REAL":                      {REAL, "C"},
	"RECURSIVE":                 {RECURSIVE, "U"},
	"REF":                       {REF, "U"},
	"REFERENCES":                {REFERENCES, "R"},
	"REGCLASS":                  {REGCLASS, "U"},
	"REGNAMESPACE":              {REGNAMESPACE, "U"},
	"REGPROC":                   {REGPROC, "U"},
	"REGPROCEDURE":              {REGPROCEDURE, "U"},
	"REGTYPE":                   {REGTYPE, "U"},
	"RELEASE":                   {RELEASE, "U"},
	"RENAME":                    {RENAME, "U"},
	"REPEATABLE":                {REPEATABLE, "U"},
	"RESET":                     {RESET, "U"},
	"RESTORE":                   {RESTORE, "U"},
	"RESTRICT":                  {RESTRICT, "U"},
	"RESUME":                    {RESUME, "U"},
	"RETURNING":                 {RETURNING, "R"},
	"REVOKE":                    {REVOKE, "U"},
	"RIGHT":                     {RIGHT, "T"},
	"ROLLBACK":                  {ROLLBACK, "U"},
	"ROLLUP":                    {ROLLUP, "U"},
	"ROW":                       {ROW, "C"},
	"ROWS":                      {ROWS, "U"},
	"SAVEPOINT":                 {SAVEPOINT, "U"},
	"SCATTER":                   {SCATTER, "U"},
	"SEARCH":                    {SEARCH, "U"},
	"SECOND":                    {SECOND, "U"},
	"SELECT":                    {SELECT, "R"},
	"SEQUENCES":                 {SEQUENCES, "U"},
	"SERIAL":                    {SERIAL, "C"},
	"SERIALIZABLE":              {SERIALIZABLE, "U"},
	"SESSION":                   {SESSION, "U"},
	"SESSIONS":                  {SESSIONS, "U"},
	"SESSION_USER":              {SESSION_USER, "R"},
	"SET":                       {SET, "U"},
	"SETTING":                   {SETTING, "U"},
	"SETTINGS":                  {SETTINGS, "U"},
	"SHOW":                      {SHOW, "U"},
	"SIMILAR":                   {SIMILAR, "T"},
	"SIMPLE":                    {SIMPLE, "U"},
	"SMALLINT":                  {SMALLINT, "C"},
	"SMALLSERIAL":               {SMALLSERIAL, "C"},
	"SNAPSHOT":                  {SNAPSHOT, "U"},
	"SOME":                      {SOME, "R"},
	"SPLIT":                     {SPLIT, "U"},
	"SQL":                       {SQL, "U"},
	"START":                     {START, "U"},
	"STATUS":                    {STATUS, "U"},
	"STDIN":                     {STDIN, "U"},
	"STORE":                     {STORE, "U"},
	"STORING":                   {STORING, "U"},
	"STRICT":                    {STRICT, "U"},
	"STRING":                    {STRING, "C"},
	"SUBSTRING":                 {SUBSTRING, "C"},
	"SYMMETRIC":                 {SYMMETRIC, "R"},
	"SYSTEM":                    {SYSTEM, "U"},
	"TABLE":                     {TABLE, "R"},
	"TABLES":                    {TABLES, "U"},
	"TEMP":                      {TEMP, "U"},
	"TEMPLATE":                  {TEMPLATE, "U"},
	"TEMPORARY":                 {TEMPORARY, "U"},
	"TESTING_RANGES":            {TESTING_RANGES, "U"},
	"TESTING_RELOCATE":          {TESTING_RELOCATE, "U"},
	"TEXT":                      {TEXT, "U"},
	"THEN":                      {THEN, "R"},
	"TIME":                      {TIME, "C"},
	"TIMESTAMP":                 {TIMESTAMP, "C"},
	"TIMESTAMPTZ":               {TIMESTAMPTZ, "C"},
	"TO":                        {TO, "R"},
	"TRACE":                     {TRACE, "U"},
	"TRAILING":                  {TRAILING, "R"},
	"TRANSACTION":               {TRANSACTION, "U"},
	"TREAT":                     {TREAT, "C"},
	"TRIM":                      {TRIM, "C"},
	"TRUE":                      {TRUE, "R"},
	"TRUNCATE":                  {TRUNCATE, "U"},
	"TYPE":                      {TYPE, "U"},
	"UNBOUNDED":                 {UNBOUNDED, "U"},
	"UNCOMMITTED":               {UNCOMMITTED, "U"},
	"UNION":                     {UNION, "R"},
	"UNIQUE":                    {UNIQUE, "R"},
	"UNKNOWN":                   {UNKNOWN, "U"},
	"UPDATE":                    {UPDATE, "U"},
	"UPSERT":                    {UPSERT, "U"},
	"USE":                       {USE, "U"},
	"USER":                      {USER, "R"},
	"USERS":                     {USERS, "U"},
	"USING":                     {USING, "R"},
	"UUID":                      {UUID, "C"},
	"VALID":                     {VALID, "U"},
	"VALIDATE":                  {VALIDATE, "U"},
	"VALUE":                     {VALUE, "U"},
	"VALUES":                    {VALUES, "C"},
	"VARCHAR":                   {VARCHAR, "C"},
	"VARIADIC":                  {VARIADIC, "R"},
	"VARYING":                   {VARYING, "U"},
	"VIEW":                      {VIEW, "R"},
	"WHEN":                      {WHEN, "R"},
	"WHERE":                     {WHERE, "R"},
	"WINDOW":                    {WINDOW, "R"},
	"WITH":                      {WITH, "R"},
	"WITHIN":                    {WITHIN, "U"},
	"WITHOUT":                   {WITHOUT, "U"},
	"WRITE":                     {WRITE, "U"},
	"YEAR":                      {YEAR, "U"},
	"ZONE":                      {ZONE, "U"},
}
