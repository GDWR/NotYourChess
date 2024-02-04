import { serverSupabaseClient } from '#supabase/server'

export default defineEventHandler(async (event) => {
    const supabase = await serverSupabaseClient(event)

    const { data } = await supabase
        .from('random_match')
        .select('*')
        .eq('active', true)
        .limit(1)
        .single();

    if (!data) {
        setResponseStatus(event, 500);
        return;
    }

    return data
}); 