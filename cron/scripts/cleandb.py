import asyncio
import asyncpg

async def main():
    print('Started cleaning db')
    conn = await asyncpg.connect('postgres://postgres@database/postgres')
    deleted_sessions = await conn.fetchval('''
        with deleted as (
            delete from sessions
            where last_used < extract(epoch from now() - '1 hour'::interval)
            returning *
        )
        select count(*) from deleted
    ''')
    print(f'Deleted {deleted_sessions} sessions')

    deleted_access_keys = await conn.fetchval('''
        with deleted as (
            delete from access_keys
            where created_at < extract(epoch from now() - '10 minutes'::interval)
            returning *
        )
        select count(*) from deleted
    ''')
    print(f'Deleted {deleted_access_keys} access keys')

if __name__ == "__main__":
    loop = asyncio.get_event_loop()
    loop.run_until_complete(main())
    loop.close()
