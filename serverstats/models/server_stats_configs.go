// Code generated by SQLBoiler (https://github.com/volatiletech/sqlboiler). DO NOT EDIT.
// This file is meant to be re-generated in place and/or deleted at any time.

package models

import (
	"context"
	"database/sql"
	"fmt"
	"reflect"
	"strconv"
	"strings"
	"sync"
	"time"

	"github.com/pkg/errors"
	"github.com/volatiletech/null"
	"github.com/volatiletech/sqlboiler/boil"
	"github.com/volatiletech/sqlboiler/queries"
	"github.com/volatiletech/sqlboiler/queries/qm"
	"github.com/volatiletech/sqlboiler/strmangle"
)

// ServerStatsConfig is an object representing the database table.
type ServerStatsConfig struct {
	GuildID        int64       `boil:"guild_id" json:"guild_id" toml:"guild_id" yaml:"guild_id"`
	CreatedAt      null.Time   `boil:"created_at" json:"created_at,omitempty" toml:"created_at" yaml:"created_at,omitempty"`
	UpdatedAt      null.Time   `boil:"updated_at" json:"updated_at,omitempty" toml:"updated_at" yaml:"updated_at,omitempty"`
	Public         null.Bool   `boil:"public" json:"public,omitempty" toml:"public" yaml:"public,omitempty"`
	IgnoreChannels null.String `boil:"ignore_channels" json:"ignore_channels,omitempty" toml:"ignore_channels" yaml:"ignore_channels,omitempty"`

	R *serverStatsConfigR `boil:"-" json:"-" toml:"-" yaml:"-"`
	L serverStatsConfigL  `boil:"-" json:"-" toml:"-" yaml:"-"`
}

var ServerStatsConfigColumns = struct {
	GuildID        string
	CreatedAt      string
	UpdatedAt      string
	Public         string
	IgnoreChannels string
}{
	GuildID:        "guild_id",
	CreatedAt:      "created_at",
	UpdatedAt:      "updated_at",
	Public:         "public",
	IgnoreChannels: "ignore_channels",
}

// ServerStatsConfigRels is where relationship names are stored.
var ServerStatsConfigRels = struct {
}{}

// serverStatsConfigR is where relationships are stored.
type serverStatsConfigR struct {
}

// NewStruct creates a new relationship struct
func (*serverStatsConfigR) NewStruct() *serverStatsConfigR {
	return &serverStatsConfigR{}
}

// serverStatsConfigL is where Load methods for each relationship are stored.
type serverStatsConfigL struct{}

var (
	serverStatsConfigColumns               = []string{"guild_id", "created_at", "updated_at", "public", "ignore_channels"}
	serverStatsConfigColumnsWithoutDefault = []string{"created_at", "updated_at", "public", "ignore_channels"}
	serverStatsConfigColumnsWithDefault    = []string{"guild_id"}
	serverStatsConfigPrimaryKeyColumns     = []string{"guild_id"}
)

type (
	// ServerStatsConfigSlice is an alias for a slice of pointers to ServerStatsConfig.
	// This should generally be used opposed to []ServerStatsConfig.
	ServerStatsConfigSlice []*ServerStatsConfig

	serverStatsConfigQuery struct {
		*queries.Query
	}
)

// Cache for insert, update and upsert
var (
	serverStatsConfigType                 = reflect.TypeOf(&ServerStatsConfig{})
	serverStatsConfigMapping              = queries.MakeStructMapping(serverStatsConfigType)
	serverStatsConfigPrimaryKeyMapping, _ = queries.BindMapping(serverStatsConfigType, serverStatsConfigMapping, serverStatsConfigPrimaryKeyColumns)
	serverStatsConfigInsertCacheMut       sync.RWMutex
	serverStatsConfigInsertCache          = make(map[string]insertCache)
	serverStatsConfigUpdateCacheMut       sync.RWMutex
	serverStatsConfigUpdateCache          = make(map[string]updateCache)
	serverStatsConfigUpsertCacheMut       sync.RWMutex
	serverStatsConfigUpsertCache          = make(map[string]insertCache)
)

var (
	// Force time package dependency for automated UpdatedAt/CreatedAt.
	_ = time.Second
)

// OneG returns a single serverStatsConfig record from the query using the global executor.
func (q serverStatsConfigQuery) OneG(ctx context.Context) (*ServerStatsConfig, error) {
	return q.One(ctx, boil.GetContextDB())
}

// One returns a single serverStatsConfig record from the query.
func (q serverStatsConfigQuery) One(ctx context.Context, exec boil.ContextExecutor) (*ServerStatsConfig, error) {
	o := &ServerStatsConfig{}

	queries.SetLimit(q.Query, 1)

	err := q.Bind(ctx, exec, o)
	if err != nil {
		if errors.Cause(err) == sql.ErrNoRows {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: failed to execute a one query for server_stats_configs")
	}

	return o, nil
}

// AllG returns all ServerStatsConfig records from the query using the global executor.
func (q serverStatsConfigQuery) AllG(ctx context.Context) (ServerStatsConfigSlice, error) {
	return q.All(ctx, boil.GetContextDB())
}

// All returns all ServerStatsConfig records from the query.
func (q serverStatsConfigQuery) All(ctx context.Context, exec boil.ContextExecutor) (ServerStatsConfigSlice, error) {
	var o []*ServerStatsConfig

	err := q.Bind(ctx, exec, &o)
	if err != nil {
		return nil, errors.Wrap(err, "models: failed to assign all query results to ServerStatsConfig slice")
	}

	return o, nil
}

// CountG returns the count of all ServerStatsConfig records in the query, and panics on error.
func (q serverStatsConfigQuery) CountG(ctx context.Context) (int64, error) {
	return q.Count(ctx, boil.GetContextDB())
}

// Count returns the count of all ServerStatsConfig records in the query.
func (q serverStatsConfigQuery) Count(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to count server_stats_configs rows")
	}

	return count, nil
}

// ExistsG checks if the row exists in the table, and panics on error.
func (q serverStatsConfigQuery) ExistsG(ctx context.Context) (bool, error) {
	return q.Exists(ctx, boil.GetContextDB())
}

// Exists checks if the row exists in the table.
func (q serverStatsConfigQuery) Exists(ctx context.Context, exec boil.ContextExecutor) (bool, error) {
	var count int64

	queries.SetCount(q.Query)
	queries.SetLimit(q.Query, 1)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return false, errors.Wrap(err, "models: failed to check if server_stats_configs exists")
	}

	return count > 0, nil
}

// ServerStatsConfigs retrieves all the records using an executor.
func ServerStatsConfigs(mods ...qm.QueryMod) serverStatsConfigQuery {
	mods = append(mods, qm.From("\"server_stats_configs\""))
	return serverStatsConfigQuery{NewQuery(mods...)}
}

// FindServerStatsConfigG retrieves a single record by ID.
func FindServerStatsConfigG(ctx context.Context, guildID int64, selectCols ...string) (*ServerStatsConfig, error) {
	return FindServerStatsConfig(ctx, boil.GetContextDB(), guildID, selectCols...)
}

// FindServerStatsConfig retrieves a single record by ID with an executor.
// If selectCols is empty Find will return all columns.
func FindServerStatsConfig(ctx context.Context, exec boil.ContextExecutor, guildID int64, selectCols ...string) (*ServerStatsConfig, error) {
	serverStatsConfigObj := &ServerStatsConfig{}

	sel := "*"
	if len(selectCols) > 0 {
		sel = strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, selectCols), ",")
	}
	query := fmt.Sprintf(
		"select %s from \"server_stats_configs\" where \"guild_id\"=$1", sel,
	)

	q := queries.Raw(query, guildID)

	err := q.Bind(ctx, exec, serverStatsConfigObj)
	if err != nil {
		if errors.Cause(err) == sql.ErrNoRows {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: unable to select from server_stats_configs")
	}

	return serverStatsConfigObj, nil
}

// InsertG a single record. See Insert for whitelist behavior description.
func (o *ServerStatsConfig) InsertG(ctx context.Context, columns boil.Columns) error {
	return o.Insert(ctx, boil.GetContextDB(), columns)
}

// Insert a single record using an executor.
// See boil.Columns.InsertColumnSet documentation to understand column list inference for inserts.
func (o *ServerStatsConfig) Insert(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) error {
	if o == nil {
		return errors.New("models: no server_stats_configs provided for insertion")
	}

	var err error
	currTime := time.Now().In(boil.GetLocation())

	if queries.MustTime(o.CreatedAt).IsZero() {
		queries.SetScanner(&o.CreatedAt, currTime)
	}
	if queries.MustTime(o.UpdatedAt).IsZero() {
		queries.SetScanner(&o.UpdatedAt, currTime)
	}

	nzDefaults := queries.NonZeroDefaultSet(serverStatsConfigColumnsWithDefault, o)

	key := makeCacheKey(columns, nzDefaults)
	serverStatsConfigInsertCacheMut.RLock()
	cache, cached := serverStatsConfigInsertCache[key]
	serverStatsConfigInsertCacheMut.RUnlock()

	if !cached {
		wl, returnColumns := columns.InsertColumnSet(
			serverStatsConfigColumns,
			serverStatsConfigColumnsWithDefault,
			serverStatsConfigColumnsWithoutDefault,
			nzDefaults,
		)

		cache.valueMapping, err = queries.BindMapping(serverStatsConfigType, serverStatsConfigMapping, wl)
		if err != nil {
			return err
		}
		cache.retMapping, err = queries.BindMapping(serverStatsConfigType, serverStatsConfigMapping, returnColumns)
		if err != nil {
			return err
		}
		if len(wl) != 0 {
			cache.query = fmt.Sprintf("INSERT INTO \"server_stats_configs\" (\"%s\") %%sVALUES (%s)%%s", strings.Join(wl, "\",\""), strmangle.Placeholders(dialect.UseIndexPlaceholders, len(wl), 1, 1))
		} else {
			cache.query = "INSERT INTO \"server_stats_configs\" %sDEFAULT VALUES%s"
		}

		var queryOutput, queryReturning string

		if len(cache.retMapping) != 0 {
			queryReturning = fmt.Sprintf(" RETURNING \"%s\"", strings.Join(returnColumns, "\",\""))
		}

		cache.query = fmt.Sprintf(cache.query, queryOutput, queryReturning)
	}

	value := reflect.Indirect(reflect.ValueOf(o))
	vals := queries.ValuesFromMapping(value, cache.valueMapping)

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, cache.query)
		fmt.Fprintln(boil.DebugWriter, vals)
	}

	if len(cache.retMapping) != 0 {
		err = exec.QueryRowContext(ctx, cache.query, vals...).Scan(queries.PtrsFromMapping(value, cache.retMapping)...)
	} else {
		_, err = exec.ExecContext(ctx, cache.query, vals...)
	}

	if err != nil {
		return errors.Wrap(err, "models: unable to insert into server_stats_configs")
	}

	if !cached {
		serverStatsConfigInsertCacheMut.Lock()
		serverStatsConfigInsertCache[key] = cache
		serverStatsConfigInsertCacheMut.Unlock()
	}

	return nil
}

// UpdateG a single ServerStatsConfig record using the global executor.
// See Update for more documentation.
func (o *ServerStatsConfig) UpdateG(ctx context.Context, columns boil.Columns) (int64, error) {
	return o.Update(ctx, boil.GetContextDB(), columns)
}

// Update uses an executor to update the ServerStatsConfig.
// See boil.Columns.UpdateColumnSet documentation to understand column list inference for updates.
// Update does not automatically update the record in case of default values. Use .Reload() to refresh the records.
func (o *ServerStatsConfig) Update(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) (int64, error) {
	currTime := time.Now().In(boil.GetLocation())

	queries.SetScanner(&o.UpdatedAt, currTime)

	var err error
	key := makeCacheKey(columns, nil)
	serverStatsConfigUpdateCacheMut.RLock()
	cache, cached := serverStatsConfigUpdateCache[key]
	serverStatsConfigUpdateCacheMut.RUnlock()

	if !cached {
		wl := columns.UpdateColumnSet(
			serverStatsConfigColumns,
			serverStatsConfigPrimaryKeyColumns,
		)

		if !columns.IsWhitelist() {
			wl = strmangle.SetComplement(wl, []string{"created_at"})
		}
		if len(wl) == 0 {
			return 0, errors.New("models: unable to update server_stats_configs, could not build whitelist")
		}

		cache.query = fmt.Sprintf("UPDATE \"server_stats_configs\" SET %s WHERE %s",
			strmangle.SetParamNames("\"", "\"", 1, wl),
			strmangle.WhereClause("\"", "\"", len(wl)+1, serverStatsConfigPrimaryKeyColumns),
		)
		cache.valueMapping, err = queries.BindMapping(serverStatsConfigType, serverStatsConfigMapping, append(wl, serverStatsConfigPrimaryKeyColumns...))
		if err != nil {
			return 0, err
		}
	}

	values := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), cache.valueMapping)

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, cache.query)
		fmt.Fprintln(boil.DebugWriter, values)
	}

	var result sql.Result
	result, err = exec.ExecContext(ctx, cache.query, values...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update server_stats_configs row")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by update for server_stats_configs")
	}

	if !cached {
		serverStatsConfigUpdateCacheMut.Lock()
		serverStatsConfigUpdateCache[key] = cache
		serverStatsConfigUpdateCacheMut.Unlock()
	}

	return rowsAff, nil
}

// UpdateAll updates all rows with the specified column values.
func (q serverStatsConfigQuery) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
	queries.SetUpdate(q.Query, cols)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update all for server_stats_configs")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to retrieve rows affected for server_stats_configs")
	}

	return rowsAff, nil
}

// UpdateAllG updates all rows with the specified column values.
func (o ServerStatsConfigSlice) UpdateAllG(ctx context.Context, cols M) (int64, error) {
	return o.UpdateAll(ctx, boil.GetContextDB(), cols)
}

// UpdateAll updates all rows with the specified column values, using an executor.
func (o ServerStatsConfigSlice) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
	ln := int64(len(o))
	if ln == 0 {
		return 0, nil
	}

	if len(cols) == 0 {
		return 0, errors.New("models: update all requires at least one column argument")
	}

	colNames := make([]string, len(cols))
	args := make([]interface{}, len(cols))

	i := 0
	for name, value := range cols {
		colNames[i] = name
		args[i] = value
		i++
	}

	// Append all of the primary key values for each column
	for _, obj := range o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), serverStatsConfigPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := fmt.Sprintf("UPDATE \"server_stats_configs\" SET %s WHERE %s",
		strmangle.SetParamNames("\"", "\"", 1, colNames),
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), len(colNames)+1, serverStatsConfigPrimaryKeyColumns, len(o)))

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, sql)
		fmt.Fprintln(boil.DebugWriter, args...)
	}

	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update all in serverStatsConfig slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to retrieve rows affected all in update all serverStatsConfig")
	}
	return rowsAff, nil
}

// UpsertG attempts an insert, and does an update or ignore on conflict.
func (o *ServerStatsConfig) UpsertG(ctx context.Context, updateOnConflict bool, conflictColumns []string, updateColumns, insertColumns boil.Columns) error {
	return o.Upsert(ctx, boil.GetContextDB(), updateOnConflict, conflictColumns, updateColumns, insertColumns)
}

// Upsert attempts an insert using an executor, and does an update or ignore on conflict.
// See boil.Columns documentation for how to properly use updateColumns and insertColumns.
func (o *ServerStatsConfig) Upsert(ctx context.Context, exec boil.ContextExecutor, updateOnConflict bool, conflictColumns []string, updateColumns, insertColumns boil.Columns) error {
	if o == nil {
		return errors.New("models: no server_stats_configs provided for upsert")
	}
	currTime := time.Now().In(boil.GetLocation())

	if queries.MustTime(o.CreatedAt).IsZero() {
		queries.SetScanner(&o.CreatedAt, currTime)
	}
	queries.SetScanner(&o.UpdatedAt, currTime)

	nzDefaults := queries.NonZeroDefaultSet(serverStatsConfigColumnsWithDefault, o)

	// Build cache key in-line uglily - mysql vs psql problems
	buf := strmangle.GetBuffer()
	if updateOnConflict {
		buf.WriteByte('t')
	} else {
		buf.WriteByte('f')
	}
	buf.WriteByte('.')
	for _, c := range conflictColumns {
		buf.WriteString(c)
	}
	buf.WriteByte('.')
	buf.WriteString(strconv.Itoa(updateColumns.Kind))
	for _, c := range updateColumns.Cols {
		buf.WriteString(c)
	}
	buf.WriteByte('.')
	buf.WriteString(strconv.Itoa(insertColumns.Kind))
	for _, c := range insertColumns.Cols {
		buf.WriteString(c)
	}
	buf.WriteByte('.')
	for _, c := range nzDefaults {
		buf.WriteString(c)
	}
	key := buf.String()
	strmangle.PutBuffer(buf)

	serverStatsConfigUpsertCacheMut.RLock()
	cache, cached := serverStatsConfigUpsertCache[key]
	serverStatsConfigUpsertCacheMut.RUnlock()

	var err error

	if !cached {
		insert, ret := insertColumns.InsertColumnSet(
			serverStatsConfigColumns,
			serverStatsConfigColumnsWithDefault,
			serverStatsConfigColumnsWithoutDefault,
			nzDefaults,
		)
		update := updateColumns.UpdateColumnSet(
			serverStatsConfigColumns,
			serverStatsConfigPrimaryKeyColumns,
		)

		if len(update) == 0 {
			return errors.New("models: unable to upsert server_stats_configs, could not build update column list")
		}

		conflict := conflictColumns
		if len(conflict) == 0 {
			conflict = make([]string, len(serverStatsConfigPrimaryKeyColumns))
			copy(conflict, serverStatsConfigPrimaryKeyColumns)
		}
		cache.query = buildUpsertQueryPostgres(dialect, "\"server_stats_configs\"", updateOnConflict, ret, update, conflict, insert)

		cache.valueMapping, err = queries.BindMapping(serverStatsConfigType, serverStatsConfigMapping, insert)
		if err != nil {
			return err
		}
		if len(ret) != 0 {
			cache.retMapping, err = queries.BindMapping(serverStatsConfigType, serverStatsConfigMapping, ret)
			if err != nil {
				return err
			}
		}
	}

	value := reflect.Indirect(reflect.ValueOf(o))
	vals := queries.ValuesFromMapping(value, cache.valueMapping)
	var returns []interface{}
	if len(cache.retMapping) != 0 {
		returns = queries.PtrsFromMapping(value, cache.retMapping)
	}

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, cache.query)
		fmt.Fprintln(boil.DebugWriter, vals)
	}

	if len(cache.retMapping) != 0 {
		err = exec.QueryRowContext(ctx, cache.query, vals...).Scan(returns...)
		if err == sql.ErrNoRows {
			err = nil // Postgres doesn't return anything when there's no update
		}
	} else {
		_, err = exec.ExecContext(ctx, cache.query, vals...)
	}
	if err != nil {
		return errors.Wrap(err, "models: unable to upsert server_stats_configs")
	}

	if !cached {
		serverStatsConfigUpsertCacheMut.Lock()
		serverStatsConfigUpsertCache[key] = cache
		serverStatsConfigUpsertCacheMut.Unlock()
	}

	return nil
}

// DeleteG deletes a single ServerStatsConfig record.
// DeleteG will match against the primary key column to find the record to delete.
func (o *ServerStatsConfig) DeleteG(ctx context.Context) (int64, error) {
	return o.Delete(ctx, boil.GetContextDB())
}

// Delete deletes a single ServerStatsConfig record with an executor.
// Delete will match against the primary key column to find the record to delete.
func (o *ServerStatsConfig) Delete(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if o == nil {
		return 0, errors.New("models: no ServerStatsConfig provided for delete")
	}

	args := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), serverStatsConfigPrimaryKeyMapping)
	sql := "DELETE FROM \"server_stats_configs\" WHERE \"guild_id\"=$1"

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, sql)
		fmt.Fprintln(boil.DebugWriter, args...)
	}

	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete from server_stats_configs")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by delete for server_stats_configs")
	}

	return rowsAff, nil
}

// DeleteAll deletes all matching rows.
func (q serverStatsConfigQuery) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if q.Query == nil {
		return 0, errors.New("models: no serverStatsConfigQuery provided for delete all")
	}

	queries.SetDelete(q.Query)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete all from server_stats_configs")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by deleteall for server_stats_configs")
	}

	return rowsAff, nil
}

// DeleteAllG deletes all rows in the slice.
func (o ServerStatsConfigSlice) DeleteAllG(ctx context.Context) (int64, error) {
	return o.DeleteAll(ctx, boil.GetContextDB())
}

// DeleteAll deletes all rows in the slice, using an executor.
func (o ServerStatsConfigSlice) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if o == nil {
		return 0, errors.New("models: no ServerStatsConfig slice provided for delete all")
	}

	if len(o) == 0 {
		return 0, nil
	}

	var args []interface{}
	for _, obj := range o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), serverStatsConfigPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "DELETE FROM \"server_stats_configs\" WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 1, serverStatsConfigPrimaryKeyColumns, len(o))

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, sql)
		fmt.Fprintln(boil.DebugWriter, args)
	}

	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete all from serverStatsConfig slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by deleteall for server_stats_configs")
	}

	return rowsAff, nil
}

// ReloadG refetches the object from the database using the primary keys.
func (o *ServerStatsConfig) ReloadG(ctx context.Context) error {
	if o == nil {
		return errors.New("models: no ServerStatsConfig provided for reload")
	}

	return o.Reload(ctx, boil.GetContextDB())
}

// Reload refetches the object from the database
// using the primary keys with an executor.
func (o *ServerStatsConfig) Reload(ctx context.Context, exec boil.ContextExecutor) error {
	ret, err := FindServerStatsConfig(ctx, exec, o.GuildID)
	if err != nil {
		return err
	}

	*o = *ret
	return nil
}

// ReloadAllG refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
func (o *ServerStatsConfigSlice) ReloadAllG(ctx context.Context) error {
	if o == nil {
		return errors.New("models: empty ServerStatsConfigSlice provided for reload all")
	}

	return o.ReloadAll(ctx, boil.GetContextDB())
}

// ReloadAll refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
func (o *ServerStatsConfigSlice) ReloadAll(ctx context.Context, exec boil.ContextExecutor) error {
	if o == nil || len(*o) == 0 {
		return nil
	}

	slice := ServerStatsConfigSlice{}
	var args []interface{}
	for _, obj := range *o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), serverStatsConfigPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "SELECT \"server_stats_configs\".* FROM \"server_stats_configs\" WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 1, serverStatsConfigPrimaryKeyColumns, len(*o))

	q := queries.Raw(sql, args...)

	err := q.Bind(ctx, exec, &slice)
	if err != nil {
		return errors.Wrap(err, "models: unable to reload all in ServerStatsConfigSlice")
	}

	*o = slice

	return nil
}

// ServerStatsConfigExistsG checks if the ServerStatsConfig row exists.
func ServerStatsConfigExistsG(ctx context.Context, guildID int64) (bool, error) {
	return ServerStatsConfigExists(ctx, boil.GetContextDB(), guildID)
}

// ServerStatsConfigExists checks if the ServerStatsConfig row exists.
func ServerStatsConfigExists(ctx context.Context, exec boil.ContextExecutor, guildID int64) (bool, error) {
	var exists bool
	sql := "select exists(select 1 from \"server_stats_configs\" where \"guild_id\"=$1 limit 1)"

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, sql)
		fmt.Fprintln(boil.DebugWriter, guildID)
	}

	row := exec.QueryRowContext(ctx, sql, guildID)

	err := row.Scan(&exists)
	if err != nil {
		return false, errors.Wrap(err, "models: unable to check if server_stats_configs exists")
	}

	return exists, nil
}
